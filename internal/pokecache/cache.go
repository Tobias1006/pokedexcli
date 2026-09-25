package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheLocation        map[string]cacheEntry
	cacheLocationPokemon map[string]cacheEntry
	mu                   sync.Mutex
}
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	var cache Cache
	cache.cacheLocation = make(map[string]cacheEntry)
	cache.cacheLocationPokemon = make(map[string]cacheEntry)
	go cache.reapLoop("loc", interval)
	go cache.reapLoop("pok", interval)
	return &cache
}

func (c *Cache) Add(ent string, key string, val []byte) {
	c.mu.Lock()
	if ent == "loc" {
		c.cacheLocation[key] = cacheEntry{
			time.Now(),
			val,
		}
	} else if ent == "pok" {
		c.cacheLocationPokemon[key] = cacheEntry{
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
	switch ent {
	case "loc":
		entry, ok := c.cacheLocation[key]
		if !ok {
			c.mu.Unlock()
			return nil, ok
		}
		c.mu.Unlock()
		val, worked = entry.val, ok
	case "pok":
		entry, ok := c.cacheLocationPokemon[key]
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
			for key, entry := range c.cacheLocation {
				if time.Since(entry.createdAt) > interval {
					c.mu.Lock()
					delete(c.cacheLocation, key)
					c.mu.Unlock()
				}
			}
		} else if ent == "pok" {
			for key, entry := range c.cacheLocationPokemon {
				if time.Since(entry.createdAt) > interval {
					c.mu.Lock()
					delete(c.cacheLocationPokemon, key)
					c.mu.Unlock()
				}
			}
		}
	}

}
