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

func BenchmarkRepeat(b *testing.B) {
	b.Run("Repeat benchmark", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Repeat("a", 5)
		}
	})
}

func ExampleRepeat_simple() {
	result := Repeat("a", 5)
	fmt.Println(result)
	// Output: aaaaa
}
