package safecounter

import (
	"sync"
)

// SafeCounter is a thread-safe counter.
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// Inc increments the counter by 1.
func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// Value returns the current value of the counter.
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
