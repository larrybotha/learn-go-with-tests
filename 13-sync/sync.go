package sync

import "sync"

type Counter struct {
	// if we pass this struct around by value, the mutex will be copied
	mu    sync.Mutex
	value int
}

// Indicate to users of the API that new Counters should be pointers
// We suggest using a pointer, because passing the Counter around by
// value will result in copies, and copies of Mutexes should be avoided
func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Value() int {
	return c.value
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++
}
