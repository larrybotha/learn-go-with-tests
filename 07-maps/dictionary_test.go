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

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("adds a value to the dictionary", func(t *testing.T) {
		key := "foo"
		value := "bar"

		dictionary := Dictionary{}
		err := dictionary.Add(key, value)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, value)
	})

	t.Run("when existing word", func(t *testing.T) {
		word := "foo"
		value := "bar"

		dictionary := Dictionary{}
		dictionary.Add(word, value)
		err := dictionary.Add(word, value)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, value)
	})
}

func assertDefinition(t *testing.T, d Dictionary, word, definition string) {
	t.Helper()

	got, err := d.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q, want %q", got, definition)
	}
}

func assertString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	if got == nil {
		if want != nil {
			t.Fatal("expected error, got none")
		}
	}
}
