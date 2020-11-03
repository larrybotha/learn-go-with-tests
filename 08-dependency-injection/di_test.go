package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("prints a value", func(t *testing.T) {
		buffer := bytes.Buffer{}

		Greet(&buffer, "Joe")

		got := buffer.String()
		want := "Hello, Joe"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
