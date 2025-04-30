package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("when word exists", func(t *testing.T) {
		value := "foo"
		dict := Dictionary{"test": value}
		key := "test"

		assertDefinition(t, dict, key, value)
	})

	t.Run("when word is unknown", func(t *testing.T) {
		dict := Dictionary{"foo": "bar"}
		key := "quux"
		_, err := dict.Search(key)

		if err == nil {
			t.Fatal("expected error")
		}

		assertErrorEq(t, err, ErrWordNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("when new word", func(t *testing.T) {
		dict := Dictionary{}
		key := "foo"
		want := "bar"

		err := dict.Add(key, want)

		assertErrorEq(t, err, nil)
		assertDefinition(t, dict, key, want)
	})

	t.Run("when word exists", func(t *testing.T) {
		key := "foo"
		value := "bar"
		dict := Dictionary{key: value}

		err := dict.Add(key, "baz")

		assertErrorEq(t, err, ErrUnableToAdd)
		assertDefinition(t, dict, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("when word exists", func(t *testing.T) {
		key := "foo"
		dict := Dictionary{key: "bar"}
		want := "baz"
		err := dict.Update(key, want)

		assertErrorEq(t, err, nil)
		assertDefinition(t, dict, key, want)
	})

	t.Run("when word is unknown", func(t *testing.T) {
		dict := Dictionary{}
		key := "foo"
		err := dict.Update(key, "bar")

		if err == nil {
			t.Fatal("expected error")
		}

		assertErrorEq(t, err, ErrUnableToUpdate)
	})
}

func TestDelete(t *testing.T) {
	t.Run("when word exists", func(t *testing.T) {
		key := "foo"
		dict := Dictionary{key: "bar"}

		err := dict.Delete(key)
		assertErrorEq(t, err, nil)

		_, err = dict.Search(key)

		assertErrorEq(t, err, ErrWordNotFound)
	})

	t.Run("when word is unknown", func(t *testing.T) {
		dict := Dictionary{}

		err := dict.Delete("foo")

		assertErrorEq(t, err, ErrUnableToDelete)
	})
}

func assertStringEq(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertErrorEq(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertDefinition(t *testing.T, dict Dictionary, key, want string) {
	t.Helper()

	got, err := dict.Search(key)
	if err != nil {
		t.Fatalf("should find expected word for key '%q'", key)
	}

	assertStringEq(t, got, want)
}
