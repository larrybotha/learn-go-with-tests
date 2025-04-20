package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("prints a value", func(t *testing.T) {
		// bytes.Buffer implements Writer
		buffer := bytes.Buffer{}

		/*
			if an address is not provided here, we get the following error:

			bytes.Buffer does not implement io.Writer (Write method has pointer receiver)

			This happens because bytes.Buffer implements Write from io.Writer with a pointer
			receiver - when an implementation of a method uses a pointer received, that value
			must be pass as a pointer
		*/
		Greet(&buffer, "Joe")

		got := buffer.String()
		want := "Hello, Joe"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
