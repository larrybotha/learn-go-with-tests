package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Sam", "")
		want := "Hello, Sam"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when an empty string is provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Sam", "Spanish")
		want := "Ola, Sam"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Sam", "French")
		want := "Bonjour, Sam"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper() // show line in tests where failed, as opposed to line where errored

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
