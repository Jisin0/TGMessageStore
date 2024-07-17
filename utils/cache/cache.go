// (c) Jisin0
//
// A minimal in-memory map cache.

package cache

import "sync"

// MapCache is a thread-safe map with a mutex.
type MapCache struct {
	mu sync.Mutex
	m  map[string]string
}

// NewMapCache initializes a new MapCache.
func NewMapCache() *MapCache {
	return &MapCache{
		m: make(map[string]string),
	}
}

// Set sets a key-value pair in the map.
func (sm *MapCache) Set(key, value string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
}

// Get retrieves a value for a given key from the map.
func (sm *MapCache) Get(key string) (string, bool) {
	value, ok := sm.m[key]
	return value, ok
}

// Delete removes a key from the map.
func (sm *MapCache) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.m, key)
}
