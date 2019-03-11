package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func producer(factor int, out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * factor
	}
}

func consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int, 64)

	go producer(3, ch)
	go producer(5, ch)

	go consumer(ch)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}
