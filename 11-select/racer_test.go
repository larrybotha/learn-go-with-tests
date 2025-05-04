package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run(
		"compares speeds of servers, returning the URL of the fastest",
		func(t *testing.T) {
			slowServer := makeDelayedServer(time.Millisecond * 10)
			fastServer := makeDelayedServer(time.Millisecond * 0)

			defer func() {
				fastServer.Close()
				slowServer.Close()
			}()

			slowUrl := slowServer.URL
			fastUrl := fastServer.URL

			want := fastUrl
			got, err := Racer(slowUrl, fastUrl)

			assertNotError(t, err)

			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		},
	)

	t.Run(
		"returns an error if a server doesn't respond within 10s",
		func(t *testing.T) {
			server := makeDelayedServer(time.Millisecond * 25)

			defer server.Close()

			_, err := ConfigurableRacer(
				server.URL,
				server.URL,
				time.Millisecond*20,
			)

			if err == nil {
				t.Error("expected error, got nil")
			}
		},
	)
}

func makeDelayedServer(n time.Duration) *httptest.Server {
	// create an http server
	return httptest.NewServer(
		// which accepts a handler
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// with a delay
			time.Sleep(n)
			// and writes to the response
			w.WriteHeader(http.StatusOK)
		}),
	)
}

func assertNotError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("got %v, wanted nil", err)
	}
}
