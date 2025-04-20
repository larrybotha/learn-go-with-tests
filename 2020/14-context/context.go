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
		// get context from request
		ctx := r.Context()
		data, err := store.Fetch(ctx)

		if err != nil {
			return
		}

		fmt.Fprint(w, data)
	}
}
