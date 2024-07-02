package eviction_policy

import (
	"testing"
)

func TestLIFOEvictionPolicy(t *testing.T) {
	capacity := 3
	policy := NewLIFOEvictionPolicy(capacity)

	keys := []string{"key1", "key2", "key3", "key4"}

	for _, key := range keys {
		policy.Access(key)
	}

	evictedKeys := make([]string, 0, capacity)
	for i := 0; i < capacity; i++ {
		key := policy.Evict()
		if key != keys[len(keys)-1-i] {
			t.Errorf("Expected evicted key %s, got %s", keys[len(keys)-1-i], key)
		}
		evictedKeys = append(evictedKeys, key)
	}

	for _, key := range keys {
		if policy.Evict() != "" {
			t.Errorf("Expected empty string for key %s after eviction", key)
		}
	}

	for _, key := range evictedKeys {
		policy.Remove(key)
	}

	if policy.stack.Len() != 0 {
		t.Error("Expected stack to be empty after removing all elements")
	}
}
