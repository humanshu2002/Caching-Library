package cache

import (
	"container/list"
)

type LIFOEvictionPolicy struct {
	stack []string
}

func NewLIFOEvictionPolicy() *LIFOEvictionPolicy {
	return &LIFOEvictionPolicy{
		stack: make([]string, 0),
	}
}

func (p *LIFOEvictionPolicy) Access(key string) *list.Element {
	p.stack = append(p.stack, key)
	return nil
}

func (p *LIFOEvictionPolicy) Evict() string {
	if len(p.stack) == 0 {
		return ""
	}
	key := p.stack[len(p.stack)-1]
	p.stack = p.stack[:len(p.stack)-1]
	return key
}

func (p *LIFOEvictionPolicy) Remove(key string) {
	for i, v := range p.stack {
		if v == key {
			p.stack = append(p.stack[:i], p.stack[i+1:]...)
			break
		}
	}
}
