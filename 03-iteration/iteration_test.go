package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleRepeat() {
	result := Repeat("x", 5)
	fmt.Println(result)
	// Output: xxxxx
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 10)
	}
}

func BenchmarkRepeatNaive(b *testing.B) {
	// legacy - prefer using b.Loop
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func TestRepeat(t *testing.T) {
	iterations := 5
	char := "a"
	result := Repeat(char, iterations)

	t.Run("contains correct number of repetitions", func(t *testing.T) {
		got := strings.Count(result, char)
		want := iterations

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
