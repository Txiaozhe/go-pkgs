package main

import (
	"log"
	"time"

	"github.com/yangwenmai/ratelimit/leakybucket"
	"github.com/yangwenmai/ratelimit/simpleratelimit"
)

func main() {
	// simple
	rl := simpleratelimit.New(1, time.Second)

	for i := 0; i < 100; i++ {
		log.Printf("limit result: %d, %v\n", i, rl.Limit())
	}
	log.Printf("limit result: %v\n", rl.Limit())

	// rate limit: leaky-bucket
	lb := leakybucket.New()
	b, err := lb.Create("leaky_bucket", 10, time.Second)
	if err != nil {
		log.Println(err)
	}
	log.Printf("bucket capacity: %v", b.Capacity())
}
