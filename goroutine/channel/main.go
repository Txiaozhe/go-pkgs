package main

import "fmt"

func oneroutine(ch chan string) {
	ch <- "this is a goroutine"
}

func main() {
	ch := make(chan string)

	go oneroutine(ch)

	select {
	case s1 := <-ch:
		fmt.Println(s1)
	}

	fmt.Println("this is main")
}
