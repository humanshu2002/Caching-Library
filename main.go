package main

import (
	"fmt"
	"time"

	"github.com/humanshu2002/Caching-Library/cache"
)

func main() {
	cacheInstance := cache.NewCache(cache.NewLRUEvictionPolicy(), 5*time.Minute)

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
