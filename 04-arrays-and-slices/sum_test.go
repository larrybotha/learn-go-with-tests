package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{5, 4, 3, 2, 1}

	t.Run("sums a slice of values", func(t *testing.T) {
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	xs := []int{1, 2, 3}
	ys := []int{4, 5, 6}

	t.Run("sums multiple slices", func(t *testing.T) {
		got := SumAll(xs, ys)
		want := []int{6, 15}

		// equality operators do not work with slices - need to use reflection
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestSumTail(t *testing.T) {
	t.Run("handles empty slices", func(t *testing.T) {
		xs := []int{}
		got := SumTail(xs)
		var want int

		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, xs)
		}
	})

	t.Run("sums only the tail of a slice", func(t *testing.T) {
		xs := []int{1, 2, 3}
		got := SumTail(xs)
		want := 5

		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, xs)
		}
	})
}

func TestSumTailAll(t *testing.T) {
	assertSums := func(t testing.TB, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("sums empty slices", func(t *testing.T) {
		xs := []int{}
		ys := []int{}

		assertSums(t, SumTailAll(xs, ys), []int{0, 0})
	})

	t.Run("sums multiple slices", func(t *testing.T) {
		xs := []int{1, 2, 3}
		ys := []int{4, 5, 6}

		assertSums(t, SumTailAll(xs, ys), []int{5, 11})
	})
}
