package eviction_policy

import (
	"container/list"
)

type LIFOEvictionPolicy struct {
	stack    *list.List
	capacity int
}

func NewLIFOEvictionPolicy(capacity int) *LIFOEvictionPolicy {
	return &LIFOEvictionPolicy{
		stack:    list.New(),
		capacity: capacity,
	}
}

func (p *LIFOEvictionPolicy) Access(key string) *list.Element {
	elem := p.stack.PushBack(key)
	if p.stack.Len() > p.capacity {
		p.stack.Remove(p.stack.Front())
	}
	return elem
}

func (p *LIFOEvictionPolicy) Evict() string {
	elem := p.stack.Back()
	if elem != nil {
		key := elem.Value.(string)
		p.stack.Remove(elem)
		return key
	}
	return ""
}

func (p *LIFOEvictionPolicy) Remove(key string) {
	for e := p.stack.Back(); e != nil; e = e.Prev() {
		if e.Value.(string) == key {
			p.stack.Remove(e)
			break
		}
	}
}
