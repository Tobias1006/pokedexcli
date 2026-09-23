package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheLocationEntries map[string]cacheEntry
	cachePokemonEntries  map[string]cacheEntry
	mu                   sync.Mutex
}
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	var cache Cache
	cache.cacheLocationEntries = make(map[string]cacheEntry)
	cache.cachePokemonEntries = make(map[string]cacheEntry)
	go cache.reapLoop("loc", interval)
	go cache.reapLoop("pok", interval)
	return &cache
}

func (c *Cache) Add(ent string, key string, val []byte) {
	c.mu.Lock()
	if ent == "loc" {
		c.cacheLocationEntries[key] = cacheEntry{
			time.Now(),
			val,
		}
	} else if ent == "pok" {
		c.cachePokemonEntries[key] = cacheEntry{
			time.Now(),
			val,
		}
	}
	c.mu.Unlock()
}

func (c *Cache) Get(ent string, key string) ([]byte, bool) {
	c.mu.Lock()
	var val []byte
	var worked bool
	if ent == "loc" {
		entry, ok := c.cacheLocationEntries[key]
		if !ok {
			c.mu.Unlock()
			return nil, ok
		}
		c.mu.Unlock()
		val, worked = entry.val, ok
	} else if ent == "pok" {
		entry, ok := c.cachePokemonEntries[key]
		if !ok {
			c.mu.Unlock()
			return nil, ok
		}
		c.mu.Unlock()
		val, worked = entry.val, ok
	}
	return val, worked
}

func (c *Cache) reapLoop(ent string, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		if ent == "loc" {
			for key, entry := range c.cacheLocationEntries {
				if time.Since(entry.createdAt) > interval {
					c.mu.Lock()
					delete(c.cacheLocationEntries, key)
					c.mu.Unlock()
				}
			}
		} else if ent == "pok" {
			for key, entry := range c.cachePokemonEntries {
				if time.Since(entry.createdAt) > interval {
					c.mu.Lock()
					delete(c.cachePokemonEntries, key)
					c.mu.Unlock()
				}
			}
		}
	}

}
