package core

import (
	"sync"
)

// Cache provides an in-memory, thread-safe cache scoped to a single discovery execution.
// It prevents redundant probing (e.g. reading /proc/meminfo or registry keys multiple times).
// It does NOT persist across runs.
type Cache interface {
	// Get retrieves a value from the cache. Returns nil, false if not found.
	Get(key string) (any, bool)

	// Set stores a value in the cache.
	Set(key string, value any)

	// Delete removes a key from the cache.
	Delete(key string)

	// Clear empties the cache.
	Clear()
}

// ephemeralCache implements Cache using a simple map and a mutex.
type ephemeralCache struct {
	mu    sync.RWMutex
	store map[string]any
}

// NewCache creates a new in-memory discovery cache.
func NewCache() Cache {
	return &ephemeralCache{
		store: make(map[string]any),
	}
}

func (c *ephemeralCache) Get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.store[key]
	return val, ok
}

func (c *ephemeralCache) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = value
}

func (c *ephemeralCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.store, key)
}

func (c *ephemeralCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store = make(map[string]any)
}
