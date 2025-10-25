package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.RWMutex{},
	}
	go newCache.reapLoop(interval)
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(s string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	data, cacheExists := c.cache[s]
	if !cacheExists {
		return nil, false
	}
	return data.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			for key, cacheValue := range c.cache {
				if time.Since(cacheValue.createdAt) > interval {
					delete(c.cache, key)
				}
			}
			c.mu.Unlock()
		}
	}
}
