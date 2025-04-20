package main

import (
	"fmt"
	"io"
	"net/http"
)

/*
By accepting a value that implements Writer

This allows for two things:

- the call site can provide whatever writer it wants
- we can now test Greet by providing a writer that doesn't write to stdout
*/
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

/*
http.ResponseWriter implements io.Writer, allowing us to pass its variable through
to Greet

Request handlers in Go take 2 arguments, similarly to Node, except that in Node
the order is switched - handler(req, res)
*/
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World\n")
}

func main() {
	/*
		In Go, one can listen and serve requests in a single call, as opposed to in Node
		where a server first needs to be created, and then listened to
	*/
	http.ListenAndServe(
		":3000",
		http.HandlerFunc(MyGreeterHandler),
	)
}
