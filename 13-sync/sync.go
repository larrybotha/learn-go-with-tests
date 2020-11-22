package sync

import "sync"

type Counter struct {
	// create a mutual exclusion lock
	// use an explicit name here - some people may define the struct with only the type
	// and no name, but this exposes Mutex as a public property on the struct, which
	// if you don't want, you shouldn't do
	mu    sync.Mutex
	count int
}

// values that contain mutexes should be passed around by reference, or as pointers.
// To encourage this, we can expose a function which creates an instance as a pointer
func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	// while this method is called, block any other goroutines from modifying this
	// instance until this function is done executing
	c.mu.Lock()
	// when this function concludes, unlock the instance. Function calls preceded by
	// `defer` are executed after all other non-deferred statements have been called
	defer c.mu.Unlock()
	c.count += 1
}

func (c *Counter) Value() int {
	return c.count
}
