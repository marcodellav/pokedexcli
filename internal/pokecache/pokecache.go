package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
type Cache struct {
	CacheEntry map[string]cacheEntry
	mu         sync.Mutex
	interval   time.Duration
}

// NewCache creates a new cache with a configurable interval
func NewCache(interval time.Duration) *Cache {
	var cache Cache
	cache.CacheEntry = make(map[string]cacheEntry)
	cache.interval = interval
	done := make(chan bool)
	go func() {
		cache.reapLoop()

		done <- true
	}()
	return &cache
}

// Add adds a new entry to the cache
func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	cache.CacheEntry[key] = entry
}

// Get gets an entry from the cache
// The returned bool should be true if the entry was found and false if it wasn't.
func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	// Use the comma-ok idiom to look up key in cache.CacheEntry
	// Return the val field of the found entry (not the whole struct) plus a bool
	// If not found, return something sensible for "no value" and false

	value, ok := cache.CacheEntry[key]
	if ok {
		return value.val, true
	} else {
		return nil, false
	}

}

// reapLoop is called when the cache is created (by the NewCache function).
// Each time an interval (the time.Duration passed to NewCache) passes
// it removes any entries that are older than the interval
func (cache *Cache) reapLoop() {

	ticker := time.NewTicker(cache.interval)
	defer ticker.Stop()

	for range ticker.C {
		cache.mu.Lock()
		for c := range cache.CacheEntry {
			if time.Since(cache.CacheEntry[c].createdAt) > cache.interval {
				delete(cache.CacheEntry, c)
			}
		}
		cache.mu.Unlock()
	}
}
