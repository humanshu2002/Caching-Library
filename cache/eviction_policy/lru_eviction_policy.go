package eviction_policy

import (
	"container/list"

	"github.com/humanshu2002/Caching-Library/structs"
)

type LRUEvictionPolicy struct {
	queue map[string]*list.Element
	list  *list.List
}

func NewLRUEvictionPolicy() *LRUEvictionPolicy {
	return &LRUEvictionPolicy{
		queue: make(map[string]*list.Element),
		list:  list.New(),
	}
}

func (p *LRUEvictionPolicy) Access(key string) *list.Element {
	if elem, found := p.queue[key]; found {
		p.list.MoveToFront(elem)
		return elem
	}
	elem := p.list.PushFront(key)
	p.queue[key] = elem
	return elem
}

func (p *LRUEvictionPolicy) Evict() string {
	elem := p.list.Back()
	if elem != nil {
		key := elem.Value.(*structs.CacheItem).Key
		delete(p.queue, key)
		p.list.Remove(elem)
		return key
	}
	return ""
}

func (p *LRUEvictionPolicy) Remove(key string) {
	if elem, found := p.queue[key]; found {
		p.list.Remove(elem)
		delete(p.queue, key)
	}
}
