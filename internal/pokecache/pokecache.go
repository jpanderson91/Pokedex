package pokecache

import (

	"sync"
	"time"
)


// cache struct (holds a map and a mutex)\
type Cache struct {
	entries map[string]cacheEntry
	interval time.Duration
	mux *sync.Mutex
}

// a cacheEntry struct (createdAt + val)
type cacheEntry struct {
	createdAt time.Time
	val []byte
}

// NewCache(interval) creates the cache and starts the reap loop
func NewCache(interval time.Duration) *Cache {
	
	c := &Cache{
		entries: make(map[string]cacheEntry),
		interval: interval,
		mux: &sync.Mutex{},
	}
	go c.reapLoop()
	return c
}


// Add(key, val) stores an entry
func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	// add an entry to the cache with the current timestamp (you can use time.Now().Unix() to get the current timestamp in seconds)
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

// Get(key) retrieves an entry
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	// get an entry from the cache, if the entry doesn't exist, return nil and false, if the entry exists, return the val and true
	entry, exists := c.entries[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

// reapLoop periodically removes stale entries
func (c *Cache) reapLoop() {
	
	// run an infinite loop that sleeps for the cache's interval (you can use time.Sleep(time.Duration(c.interval) * time.Second)), after waking up, it should check all entries in the cache and remove any entries that are older than the cache's interval
	for {
		time.Sleep(c.interval)
		c.mux.Lock()
		now := time.Now()
		for key, entry := range c.entries {
			if entry.createdAt.Before(now.Add(-c.interval)) {
				delete(c.entries, key)
			}
		}
		c.mux.Unlock()
	}
}

