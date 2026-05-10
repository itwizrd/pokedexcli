package pokecache

import (
	"time"
)

func NewCache(interval time.Duration) Cache {
	return Cache{
		cacheEntry: cacheEntry{
			createdAt: interval
			val: []byte
		}
	}
}

//methods
func (c *Cache) Add(key string, val []byte) {
	//add string to cache map
}

func (c *Cache) Get(key string) ([]byte, bool) {
	//use string to lookup cache map
	//return true if found, false if not.
	//return the []byte
}

func reapLoop(interval time.Duration) {
	// called when cache is created by NewCache()
	// remove any entries older than the interval
	// use time.Ticker
}