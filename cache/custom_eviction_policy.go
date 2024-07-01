package cache

import (
	"container/list"
	"time"
)

type ItemInfo struct {
	key        string
	lastAccess time.Time
	frequency  int
}

// Frequency and Recency-based Eviction Policy" (FREP)
type FREPEvictionPolicy struct {
	capacity int
	items    map[string]*list.Element
	order    *list.List
}

func NewCustomEvictionPolicy(capacity int) *FREPEvictionPolicy {
	return &FREPEvictionPolicy{
		capacity: capacity,
		items:    make(map[string]*list.Element),
		order:    list.New(),
	}
}

func (p *FREPEvictionPolicy) Access(key string) *list.Element {
	if element, exists := p.items[key]; exists {
		info := element.Value.(*ItemInfo)
		info.lastAccess = time.Now()
		info.frequency++
		p.order.MoveToFront(element)
		return element
	}
	return nil
}

func (p *FREPEvictionPolicy) Evict() string {
	if p.order.Len() == 0 {
		return ""
	}

	var evictElement *list.Element
	for e := p.order.Back(); e != nil; e = e.Prev() {
		if evictElement == nil {
			evictElement = e
		} else {
			eInfo := e.Value.(*ItemInfo)
			evictInfo := evictElement.Value.(*ItemInfo)

			if eInfo.frequency < evictInfo.frequency ||
				(eInfo.frequency == evictInfo.frequency && eInfo.lastAccess.Before(evictInfo.lastAccess)) {
				evictElement = e
			}
		}
	}

	if evictElement != nil {
		evictInfo := evictElement.Value.(*ItemInfo)
		p.order.Remove(evictElement)
		delete(p.items, evictInfo.key)
		return evictInfo.key
	}

	return ""
}

func (p *FREPEvictionPolicy) Remove(key string) {
	if element, exists := p.items[key]; exists {
		p.order.Remove(element)
		delete(p.items, key)
	}
}

func (p *FREPEvictionPolicy) Add(key string) {
	if p.order.Len() >= p.capacity {
		p.Evict()
	}
	itemInfo := &ItemInfo{
		key:        key,
		lastAccess: time.Now(),
		frequency:  1,
	}
	element := p.order.PushFront(itemInfo)
	p.items[key] = element
}
