package main

import (
	"fmt"
	"time"

	"github.com/humanshu2002/Caching-Library/cache"
	"github.com/humanshu2002/Caching-Library/cache/eviction_policy"
	"github.com/humanshu2002/Caching-Library/structs"
)

func testCache(evictionPolicy structs.EvictionPolicy, cacheName string) {
	fmt.Println("Testing", cacheName)
	cacheInstance := cache.NewCache(evictionPolicy, 5*time.Minute, 5)

	cacheInstance.Set("key1", "value1", 5*time.Minute)
	value, err := cacheInstance.Get("key1")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Value:", value)
	}

	cacheInstance.Delete("key1")
	value, err = cacheInstance.Get("key1")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Value:", value)
	}
}

func main() {
	testCache(eviction_policy.NewLRUEvictionPolicy(), "LRU Eviction Policy")
	testCache(eviction_policy.NewFIFOEvictionPolicy(), "FIFO Eviction Policy")
	testCache(eviction_policy.NewLIFOEvictionPolicy(5), "LIFO Eviction Policy")
	testCache(eviction_policy.NewCustomEvictionPolicy(10), "Custom Eviction Policy")
}
