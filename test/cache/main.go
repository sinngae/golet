package main

import (
	"time"

	"github.com/patrickmn/go-cache"
)

func main() {
	myCache := cache.Cache{}
	myCache.Set("test", "abc", 3*time.Second)
}
