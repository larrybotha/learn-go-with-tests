package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

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
		data := "foo"
		store := &StoreSpy{response: data, t: t}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingContext, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingContext)

		response := &ResponseWriterSpy{}

		server.ServeHTTP(response, request)

		fmt.Printf("%v\n", response)

		if response.written {
			t.Error("no response should have been written")
		}
	})
}

// define our own ResponseWriter spy because httptest.ResponseRecorder doesn't allow
// for us to evaluate whether a response is written or not
type ResponseWriterSpy struct {
	written bool
}

// ResponseWrite has a Header method we need to implement
func (r *ResponseWriterSpy) Header() http.Header {
	r.written = true
	return nil
}

// ResponseWriter has a Write  method we need to implement
func (r *ResponseWriterSpy) Write([]byte) (int, error) {
	r.written = true
	return 0, errors.New("not implemented")
}

// ResponseWriter has a WriteHeader method we need to implement
func (r *ResponseWriterSpy) WriteHeader(statusCode int) {
	r.written = true
}

type StoreSpy struct {
	response  string
	cancelled bool
	t         *testing.T
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

func (s *StoreSpy) Fetch(ctx context.Context) (string, error) {
	// create a channel that will receive 1 value
	data := make(chan string, 1)

	// start a goroutine
	go func() {
		var result string

		// iterate over the response property of the store
		// we iterate over individual characters simulate a long request
		// - this seems odd... why not simply iterate over a slice of length 1 with
		// a long timeout?
		for _, c := range s.response {
			// select the first channel to receive a value
			select {
			// if the context is done, we know that it's been cancelled, either from
			// an explicit cancel, or because of an error
			case <-ctx.Done():
				s.t.Log("spy store got cancelled")
				return

				// otherwise we simulate a request taking some time
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}

		// once we are done iterating over s.response, send the result to the channel
		data <- result
	}()

	// this select waits for the goroutine to complete
	select {
	// if "cancel" is called for the context, return an empty string, and return
	// the reason for the cancellation
	case <-ctx.Done():
		return "", ctx.Err()
		// if our channel receives a value, then return that value, and a nil error
	case res := <-data:
		return res, nil
	}
}

func (s *StoreSpy) Cancel() {
	s.cancelled = true
}
