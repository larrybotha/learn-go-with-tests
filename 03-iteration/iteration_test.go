package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeats a value", func(t *testing.T) {
		got := Repeat("a", 5)
		expected := "aaaaa"

		if got != expected {
			t.Errorf("expected %q to be %q", got, expected)
		}
	})
}

func ExampleRepeat_simple() {
	result := Repeat("a", 5)
	fmt.Println(result)
	// Output: aaaaa
}
