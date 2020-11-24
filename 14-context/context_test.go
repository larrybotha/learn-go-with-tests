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
		server := Server(&StoreSpy{response: data})

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		body := response.Body.String()

		if body != data {
			t.Errorf("got %q, want %q", body, data)
		}
	})

	t.Run("allows cancelling requests", func(t *testing.T) {
		store := &StoreSpy{response: "foo"}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingContext, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingContext)

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if !store.cancelled {
			t.Errorf("got %v, want %v", store.cancelled, true)
		}
	})
}
