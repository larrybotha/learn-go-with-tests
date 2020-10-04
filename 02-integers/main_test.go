package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("adds numbers", func(t *testing.T) {
		got := Add(2, 3)
		want := 5

		if got != want {
			t.Errorf("expected %d to be %d", got, want)
		}
	})
}

func ExampleAdd_simple() {
	sum := Add(2, 3)
	fmt.Println(sum)
	// Output: 5
}
