package cache

import "time"

type CacheItem struct {
	Key      string
	Value    interface{}
	ExpireAt time.Time
}
