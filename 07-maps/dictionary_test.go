package main

import (
	"testing"
)

func TestSearch_naive(t *testing.T) {
	t.Run("returns a value at a location", func(t *testing.T) {
		value := "this is just a test"
		dictionary := map[string]string{"test": value}

		got := Search(dictionary, "test")
		want := value

		assertString(t, got, want)
	})
}

func TestSearch(t *testing.T) {
	t.Run("gets value from map", func(t *testing.T) {
		value := "test string"
		dictionary := Dictionary{"test": value}

		got := dictionary.Search("test")
		want := value

		assertString(t, got, want)
	})
}

func assertString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
