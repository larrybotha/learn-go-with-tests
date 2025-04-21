package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(2, 5)
	fmt.Println(sum)
	// Output: 7
}

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	want := 4

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
