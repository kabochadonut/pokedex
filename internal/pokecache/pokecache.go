package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	ma map[string]CacheEntry
	mu *sync.RWMutex
}

type CacheEntry struct {
	created_at time.Time
	val        []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		ma: make(map[string]CacheEntry),
		mu: &sync.RWMutex{},
	}
	go c.reapLoop(time.Duration(interval))
	return c
}

func (c *Cache) Add(name string, val []byte) {
	c.mu.Lock()
	c.ma[name] = CacheEntry{
		created_at: time.Now(),
		val:        val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(name string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if entry, ok := c.ma[name]; !ok {
		return []byte{}, false
	} else {
		return entry.val, true
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		for key, entry := range c.ma {
			if time.Now().After(entry.created_at.Add(interval)) {
				c.mu.Lock()
				delete(c.ma, key)
				c.mu.Unlock()
			}
		}
	}
}
