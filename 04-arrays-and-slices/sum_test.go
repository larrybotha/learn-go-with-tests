package sum

import (
	"fmt"
	"testing"
)

func TestSum5(t *testing.T) {
	t.Run("returns the sum of 5 integers in an array", func(t *testing.T) {
		// xs := [...]int{1, 2, 3, 4, 5}
		xs := [5]int{1, 2, 3, 4, 5}

		actual := Sum5(xs)
		expected := 15

		if actual != expected {
			t.Errorf("expected %d to be %d, %v", actual, expected, xs)
		}
	})
}

func TestSum(t *testing.T) {
	t.Run("returns the sum of integers in a slice", func(t *testing.T) {
		xs := []int{1, 2, 3, 4, 5}

		actual := Sum(xs)
		expected := 15

		if actual != expected {
			t.Errorf("expected %d to be %d, %v", actual, expected, xs)
		}
	})
}

func ExampleSum() {
	result := Sum([]int{1, 2, 3})
	fmt.Println(result)
	// Output: 6
}
