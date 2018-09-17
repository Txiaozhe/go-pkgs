package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r1 := rand.Intn(10)
	fmt.Println("rand1: ", r1) // 不设置种子时随机数一样

	rand.Seed(time.Now().Unix())
	r2 := rand.Intn(10)
	fmt.Println("rand2: ", r2)
}
