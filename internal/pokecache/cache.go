package pokecache

import (
	"time"
	"sync"
)

func NewCache(interval time.Duration) Cache {
	m := map[string]cacheEntry{}

	mux := &sync.Mutex{}

	//go


	return Cache{
		cacheEntry: cacheEntry{
			createdAt: interval,
			val: []byte,
		}
	}
}

//methods
func (c *Cache) Add(key string, val []byte) {
	//add string to cache map
	c.m[key] = val
}

func (c *Cache) Get(key string) ([]byte, bool) {
	//use string to lookup cache map
	//return true if found, false if not.
	//return the []byte
	b, ok := c.m[key]
	if ok {
		return b.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration, mux *sync.Mutex) {
	// called when cache is created by NewCache()
	// remove any entries older than the interval
	// use time.Ticker
}
