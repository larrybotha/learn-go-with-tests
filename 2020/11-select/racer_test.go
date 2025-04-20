package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRace(t *testing.T) {
	t.Run("returns the faster of 2 requests", func(t *testing.T) {
		slowServer := createDelayedServer(20 * time.Millisecond)
		fastServer := createDelayedServer(0 * time.Millisecond)

		/*
			`defer` calls a function at the end of its containing function
		*/
		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, err := Racer(slowUrl, fastUrl)

		/*
			prevent further execution of this test by using t.Fatal if we get an error
		*/
		if err != nil {
			t.Fatalf("got an error, didn't expect one: %v", err)
		}

		if got != want {
			t.Errorf("go %q, want %q", got, want)
		}
	})

	t.Run("returns an error if requests take longer than 10 seconds", func(t *testing.T) {
		slowServer := createDelayedServer(12 * time.Millisecond)
		fastServer := createDelayedServer(11 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		_, err := ConfigurableRacer(slowServer.URL, fastServer.URL, 13*time.Millisecond)

		if err == nil {
			t.Error("expected error")
		}
	})
}

func createDelayedServer(duration time.Duration) *httptest.Server {
	handler := createDelayedHandler(duration)
	server := httptest.NewServer(http.HandlerFunc(handler))

	return server
}

func createDelayedHandler(duration time.Duration) http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}

	return handler
}
