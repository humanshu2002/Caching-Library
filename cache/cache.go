package cache

import (
	"container/list"
	"errors"
	"sync"
	"time"

	"github.com/humanshu2002/Caching-Library/structs"
)

type Cache struct {
	evictionPolicy structs.EvictionPolicy
	items          map[string]*list.Element
	capacity       int
	mu             sync.Mutex
	ttl            time.Duration
}

func NewCache(evictionPolicy structs.EvictionPolicy, ttl time.Duration, capacity int) *Cache {
	return &Cache{
		evictionPolicy: evictionPolicy,
		items:          make(map[string]*list.Element),
		capacity:       capacity,
		ttl:            ttl,
	}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.items[key]; ok {
		c.evictionPolicy.Access(key)
		cacheItem := elem.Value.(*structs.CacheItem)
		cacheItem.Value = value
		cacheItem.Expiration = time.Now().Add(ttl).UnixNano()
		return
	}

	if len(c.items) >= c.capacity {
		evictedKey := c.evictionPolicy.Evict()
		if evictedKey != "" {
			delete(c.items, evictedKey)
		}
	}

	cacheItem := &structs.CacheItem{
		Key:        key,
		Value:      value,
		Expiration: time.Now().Add(ttl).UnixNano(),
	}
	elem := c.evictionPolicy.Access(key)
	c.items[key] = elem
	elem.Value = cacheItem
}

func (c *Cache) Get(key string) (interface{}, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.items[key]; ok {
		cacheItem := elem.Value.(*structs.CacheItem)
		if cacheItem.IsExpired() {
			c.evictionPolicy.Remove(key)
			delete(c.items, key)
			return nil, errors.New("item has expired")
		}
		c.evictionPolicy.Access(key)
		return cacheItem.Value, nil
	}
	return nil, errors.New("item not found")
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.items[key]; ok {
		c.evictionPolicy.Remove(key)
		delete(c.items, key)
	}
}
