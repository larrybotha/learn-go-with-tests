package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run(
		"incrementing the counter 3 times leaves it at 3",
		func(t *testing.T) {
			counter := NewCounter()
			increments := 3

			for range increments {
				counter.Inc()
			}

			assertCounter(t, counter, increments)
		},
	)

	t.Run("runs safely concurrencly", func(t *testing.T) {
		counter := NewCounter()
		increments := 1000

		var wg sync.WaitGroup

		wg.Add(increments) // specify the number of goroutines to wait for

		for range increments {
			go func() {
				counter.Inc()
				wg.Done() // decrement the waitGroup
			}()
		}

		wg.Wait() // block until the items in the group are done

		assertCounter(t, counter, increments)
	})
}

// pass a pointer to the counter through, otherwise a copy of the mutex
// is created
func assertCounter(t *testing.T, c *Counter, want int) {
	t.Helper()

	if c.Value() != want {
		t.Errorf("got %d, want %d", c.Value(), want)
	}
}
