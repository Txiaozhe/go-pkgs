package main

import (
	"fmt"

	"github.com/spaolacci/murmur3"
)

var bucketSize uint64 = 10

// 15810000000 开头造一千万个手机号计算哈希值并分布在10个bucket
func main() {
	var bucketMap = map[uint64]int{}
	for i := 15810000000; i < 15810000000+10000000; i++ {
		hashInt := murmur64(fmt.Sprint(i)) % bucketSize
		bucketMap[hashInt]++
	}

	fmt.Println(bucketMap)
}

func murmur64(p string) uint64 {
	return murmur3.Sum64([]byte(p))
}
