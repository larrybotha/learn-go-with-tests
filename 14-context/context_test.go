package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type StoreSpy struct {
	response  string
	cancelled bool
	t         *testing.T
}

func (s *StoreSpy) Fetch() string {
	time.Sleep(10 * time.Millisecond)
	return s.response
}

func (s *StoreSpy) Cancel() {
	s.cancelled = true
}

func TestHandler(t *testing.T) {
	t.Run("handles basic requests", func(t *testing.T) {
		data := "hello, world"
		store := &StoreSpy{response: data, t: t}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		body := response.Body.String()

		if body != data {
			t.Errorf("got %q, want %q", body, data)
		}

		store.assertWasNotCancelled()
	})

	t.Run("allows cancelling requests", func(t *testing.T) {
		store := &StoreSpy{response: "foo", t: t}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		// Create a derived context from the request
		// context.WithCancel returns a cancel function which can cancel the context
		cancellingContext, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		// create a shallow copy of request, setting the cancellingContext as its
		// context
		request = request.WithContext(cancellingContext)

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		store.assertWasCancelled()
	})
}

func (s *StoreSpy) assertWasCancelled() {
	s.t.Helper()

	if !s.cancelled {
		s.t.Error("was not cancelled when it should have been")
	}
}

func (s *StoreSpy) assertWasNotCancelled() {
	s.t.Helper()

	if s.cancelled {
		s.t.Error("was cancelled when it should not have been")
	}
}
