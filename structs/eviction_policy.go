package structs

import "container/list"

type EvictionPolicy interface {
	Access(key string) *list.Element
	Evict() string
	Remove(key string)
}
