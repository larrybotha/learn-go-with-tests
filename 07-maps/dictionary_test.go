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
	t.Run("when known word", func(t *testing.T) {
		value := "test string"
		dictionary := Dictionary{"test": value}

		got, _ := dictionary.Search("test")
		want := value

		assertString(t, got, want)
	})

	t.Run("when unknown word", func(t *testing.T) {
		value := "test string"
		dictionary := Dictionary{"test": value}

		_, err := dictionary.Search("foo")

		assertError(t, err, ErrNoValue)
	})
}

func TestAdd(t *testing.T) {
	t.Run("adds a value to the dictionary", func(t *testing.T) {
		key := "foo"
		value := "bar"

		dictionary := Dictionary{}
		dictionary.Add(key, value)

		got, _ := dictionary.Search(key)
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

func assertError(t *testing.T, err error, want error) {
	t.Helper()

	if err == nil {
		t.Fatal("expected error, got none")
	}

	if err != want {
		t.Errorf("got %q, want %q", err, want)
	}
}
