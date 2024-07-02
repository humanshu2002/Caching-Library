package cache

import (
	"testing"
	"time"

	"github.com/humanshu2002/Caching-Library/cache/eviction_policy"
)

func TestCacheWithLIFO(t *testing.T) {
	cache := NewCache(eviction_policy.NewLIFOEvictionPolicy(3), 5*time.Minute, 3)

	cache.Set("key1", "value1", 5*time.Minute)
	cache.Set("key2", "value2", 5*time.Minute)
	cache.Set("key3", "value3", 5*time.Minute)

	if val, err := cache.Get("key1"); err != nil || val != "value1" {
		t.Errorf("Expected value1, got %s", val)
	}

	cache.Set("key4", "value4", 5*time.Minute)
	if _, err := cache.Get("key1"); err == nil {
		t.Error("Expected error for key1 as it should be evicted")
	}
}

func TestCacheWithFIFO(t *testing.T) {
	cache := NewCache(eviction_policy.NewFIFOEvictionPolicy(), 5*time.Minute, 3)

	cache.Set("key1", "value1", 5*time.Minute)
	cache.Set("key2", "value2", 5*time.Minute)
	cache.Set("key3", "value3", 5*time.Minute)

	if val, err := cache.Get("key1"); err != nil || val != "value1" {
		t.Errorf("Expected value1, got %s", val)
	}

	cache.Set("key4", "value4", 5*time.Minute)
	if _, err := cache.Get("key1"); err == nil {
		t.Error("Expected error for key1 as it should be evicted")
	}
}

func TestCacheWithLRU(t *testing.T) {
	cache := NewCache(eviction_policy.NewLRUEvictionPolicy(), 5*time.Minute, 3)

	cache.Set("key1", "value1", 5*time.Minute)
	cache.Set("key2", "value2", 5*time.Minute)
	cache.Set("key3", "value3", 5*time.Minute)

	if val, err := cache.Get("key1"); err != nil || val != "value1" {
		t.Errorf("Expected value1, got %s", val)
	}

	cache.Set("key4", "value4", 5*time.Minute)
	if _, err := cache.Get("key2"); err == nil {
		t.Error("Expected error for key2 as it should be evicted")
	}
}
