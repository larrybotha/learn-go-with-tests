package server

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*
		// get context from request
		ctx := r.Context()
		// create a channel which will receive a string, and has size 1
		data := make(chan string, 1)

		// fire a goroutine to fetch the data asynchronously
		go func() {
			// send the result of store.Fetch() to the channel
			data <- store.Fetch()
		}()

		// race the asynchronous processes
		select {
		// if our channel receives a value before the context is Done or Cancelled,
		// send the received value to the response writer
		case d := <-data:
			fmt.Fprintf(w, d)
		// If the context is cancelled or is done, cancel the request.
		// With Cancel being defined on Store's interface, and store being passed into
		// Server, Cancel essentially behaves like a callback - the call-site is
		// responsible for doing what it wants when Cancel is executed
		case <-ctx.Done():
			store.Cancel()
			*/
		}
	}
}
