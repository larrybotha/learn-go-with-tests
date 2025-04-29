package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// Buffer implements io.Writer because it has a .Write method
	buffer := bytes.Buffer{}

	Greet(&buffer, "Sam")

	want := "Hello, Sam"
	got := buffer.String()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
