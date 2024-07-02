package eviction_policy

import (
	"container/list"

	"github.com/humanshu2002/Caching-Library/structs"
)

type CustomEvictionPolicy struct {
	capacity int
	items    map[string]*list.Element
	order    *list.List
}

func NewCustomEvictionPolicy(capacity int) *CustomEvictionPolicy {
	return &CustomEvictionPolicy{
		capacity: capacity,
		items:    make(map[string]*list.Element),
		order:    list.New(),
	}
}

func (p *CustomEvictionPolicy) Access(key string) *list.Element {
	if elem, found := p.items[key]; found {
		p.order.MoveToFront(elem)
		return elem
	}
	elem := p.order.PushFront(key)
	p.items[key] = elem
	return elem
}

func (p *CustomEvictionPolicy) Evict() string {
	elem := p.order.Back()
	if elem != nil {
		key := elem.Value.(*structs.CacheItem).Key
		delete(p.items, key)
		p.order.Remove(elem)
		return key
	}
	return ""
}

func (p *CustomEvictionPolicy) Remove(key string) {
	if elem, found := p.items[key]; found {
		p.order.Remove(elem)
		delete(p.items, key)
	}
}
