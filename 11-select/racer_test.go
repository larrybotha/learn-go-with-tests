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
			got := Racer(slowUrl, fastUrl)

			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		},
	)

	t.Run(
		"returns an error if a server doesn't respond within 10s",
		func(t *testing.T) {
			slowServer := makeDelayedServer(time.Second * 12)
			fastServer := makeDelayedServer(time.Second * 11)

			defer func() {
				fastServer.Close()
				slowServer.Close()
			}()

			slowUrl := slowServer.URL
			fastUrl := fastServer.URL

			want := fastUrl
			got := Racer(slowUrl, fastUrl)

			if got != want {
				t.Errorf("got %q, want %q", got, want)
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
