package db

import (
	"sync"
	"time"
)

type entry struct {
	value      interface{}
	expiration time.Time
}

type CacheDB struct {
	capacity int
	items    map[string]entry
	order    []string
	mu       sync.Mutex
}

func NewCacheDB(capacity int) *CacheDB {
	return &CacheDB{
		capacity: capacity,
		items:    make(map[string]entry),
		order:    make([]string, 0),
	}
}

func (c *CacheDB) Set(key string, value interface{}, duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	expiration := time.Now().Add(duration)
	c.items[key] = entry{value, expiration}
	c.updateOrder(key)

	if len(c.items) > c.capacity {
		c.evictOldest()
	}
}

func (c *CacheDB) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if item, exists := c.items[key]; exists {
		if item.expiration.After(time.Now()) {
			c.updateOrder(key)
			return item.value, true
		} else {
			c.delete(key)
		}
	}
	return nil, false
}

func (c *CacheDB) updateOrder(key string) {
	for i, k := range c.order {
		if k == key {
			c.order = append(c.order[:i], c.order[i+1:]...)
			break
		}
	}
	c.order = append([]string{key}, c.order...)
}

func (c *CacheDB) evictOldest() {
	if len(c.order) == 0 {
		return
	}
	oldestKey := c.order[len(c.order)-1]
	c.delete(oldestKey)
}

func (c *CacheDB) delete(key string) {
	delete(c.items, key)
	for i, k := range c.order {
		if k == key {
			c.order = append(c.order[:i], c.order[i+1:]...)
			break
		}
	}
}
