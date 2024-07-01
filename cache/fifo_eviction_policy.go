package cache

import (
	"container/list"
)

type FIFOEvictionPolicy struct {
	queue *list.List
}

func NewFIFOEvictionPolicy() *FIFOEvictionPolicy {
	return &FIFOEvictionPolicy{
		queue: list.New(),
	}
}

func (p *FIFOEvictionPolicy) Access(key string) *list.Element {
	return p.queue.PushBack(key)
}

func (p *FIFOEvictionPolicy) Evict() string {
	elem := p.queue.Front()
	if elem != nil {
		key := elem.Value.(string)
		p.queue.Remove(elem)
		return key
	}
	return ""
}

func (p *FIFOEvictionPolicy) Remove(key string) {
	for e := p.queue.Front(); e != nil; e = e.Next() {
		if e.Value.(string) == key {
			p.queue.Remove(e)
			break
		}
	}
}
