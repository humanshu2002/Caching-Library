package eviction_policy

import (
	"testing"
	"time"

	"github.com/humanshu2002/Caching-Library/cache"
	"github.com/stretchr/testify/assert"
)

func TestLRUEvictionPolicy(t *testing.T) {
	cacheInstance := cache.NewCache(NewLRUEvictionPolicy(), 5*time.Minute, 5)
	cacheInstance.Set("key1", "value1", 5*time.Minute)
	cacheInstance.Set("key2", "value2", 5*time.Minute)
	cacheInstance.Set("key3", "value3", 5*time.Minute)
	cacheInstance.Set("key4", "value4", 5*time.Minute)
	cacheInstance.Set("key5", "value5", 5*time.Minute)
	cacheInstance.Set("key6", "value6", 5*time.Minute)

	_, err := cacheInstance.Get("key1")
	assert.Error(t, err)

	value, err := cacheInstance.Get("key6")
	assert.NoError(t, err)
	assert.Equal(t, "value6", value)
}
