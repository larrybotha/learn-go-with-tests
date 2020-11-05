package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("prints a value", func(t *testing.T) {
		// bytes.Buffer implements Writer
		buffer := bytes.Buffer{}

		// why is buffer passed in with an &?
		Greet(&buffer, "Joe")

		got := buffer.String()
		want := "Hello, Joe"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
