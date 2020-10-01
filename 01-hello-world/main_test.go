package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectString := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("name")
		want := "Hello, name"

		assertCorrectString(t, got, want)
	})

	t.Run("saying hello with default", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectString(t, got, want)
	})
}
