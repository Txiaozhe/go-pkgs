package main

import (
	"fmt"
)

func GenerateNumber() chan int {
	ch := make(chan int)

	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()

	return ch
}

func PrimerFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()

	return out
}

func main() {
	ch := GenerateNumber()
	for i := 0; i < 25; i++ {
		prime := <-ch
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimerFilter(ch, prime)
	}
}
