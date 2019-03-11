package main

import (
	"fmt"
	"time"

	cache "github.com/patrickmn/go-cache"
)

func main() {
	c := cache.New(5*time.Minute, 10*time.Minute)
	c.Set("foo", "bar", cache.DefaultExpiration)
	c.Set("baz", 42, cache.DefaultExpiration)

	foo, found := c.Get("foo")
	if found {
		// 此处foo需断言
		fmt.Println(foo.(string))
	}
}
