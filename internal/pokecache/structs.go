package pokecache

import(
	"time"
)
type Locations struct
type Cache struct {
	map[string]cacheEntry
	//sync.Mutex that projects map across goroutines
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}