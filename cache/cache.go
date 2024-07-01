package cache

import (
	"container/list"
	"errors"
	"sync"
	"time"
)

type Cache interface {
	Set(key string, value interface{}, ttl time.Duration)
	Get(key string) (interface{}, error)
	Delete(key string)
}

type CacheImpl struct {
	items          map[string]*list.Element
	evictionPolicy EvictionPolicy
	mu             sync.Mutex
	ttl            time.Duration
}

func NewCache(policy EvictionPolicy, ttl time.Duration) *CacheImpl {
	return &CacheImpl{
		items:          make(map[string]*list.Element),
		evictionPolicy: policy,
		ttl:            ttl,
	}
}

func (c *CacheImpl) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, found := c.items[key]; found {
		c.evictionPolicy.Remove(key)
		c.evictionPolicy.Access(key)
		elem.Value.(*CacheItem).Value = value
		elem.Value.(*CacheItem).ExpireAt = time.Now().Add(ttl)
		return
	}

	item := &CacheItem{
		Key:      key,
		Value:    value,
		ExpireAt: time.Now().Add(ttl),
	}
	elem := c.evictionPolicy.Access(key)
	c.items[key] = elem
	elem.Value = item
}

func (c *CacheImpl) Get(key string) (interface{}, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, found := c.items[key]; found {
		item := elem.Value.(*CacheItem)
		if item.ExpireAt.After(time.Now()) {
			c.evictionPolicy.Access(key)
			return item.Value, nil
		}
		c.Delete(key)
	}
	return nil, errors.New("key not found")
}

func (c *CacheImpl) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, found := c.items[key]; found {
		c.evictionPolicy.Remove(key)
		delete(c.items, key)
	}
}
