package structs

import "time"

type CacheItem struct {
	Key        string
	Value      interface{}
	Expiration int64
}

func (item *CacheItem) IsExpired() bool {
	if item.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > item.Expiration
}
