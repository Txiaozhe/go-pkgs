package main

import (
	"fmt"
	"time"
)

func oneroutine() {
	fmt.Println("this is a goroutine")
}

func main() {
	go oneroutine()
	time.Sleep(time.Second)
	fmt.Println("this is main")
}
