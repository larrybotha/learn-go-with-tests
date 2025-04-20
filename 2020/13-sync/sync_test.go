package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter increases the count", func(t *testing.T) {
		counter := NewCounter()

		counter.Inc()
		counter.Inc()
		counter.Inc()

		want := 3

		assertCounter(t, counter, want)
	})

	t.Run("runs safely when concurrent", func(t *testing.T) {
		want := 1000
		counter := NewCounter()

		// sync.WaitGroup can be used to synchronise concurrent processes
		var wg sync.WaitGroup
		wg.Add(want)

		for i := 0; i < want; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}

		wg.Wait()

		assertCounter(t, counter, want)
	})
}

// assertCounter requires a pointer to the counter to be passed in, because
// counter contains a mutex. A mutex may only be copied after first use, and to
// avoid further copying, the counter needs to be passed around by reference, and
// not as a copy
func assertCounter(t *testing.T, c *Counter, want int) {
	t.Helper()

	got := c.Value()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
