package main

import (
	"context"
	"fmt"
	"time"
)

func inc(a int) int {
	res := a + 1
	time.Sleep(1 * time.Second)
	return res
}

func Add(ctx context.Context, a, b int) int {
	res := 0
	for i := 0; i < a; i++ {
		res = inc(res)
		select {
		case <-ctx.Done():
			return -1
		default:
		}
	}

	for j := 0; j < b; j++ {
		res = inc(res)
		select {
		case <-ctx.Done():
			return -1
		default:
		}
	}

	return res
}

func main() {
	{
		a := 1
		b := 2
		timeout := 2 * time.Second
		ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
		res := Add(ctx, 1, 2)
		cancelFunc() // need to call
		fmt.Printf("Compute: %d+%d, result: %d\n", a, b, res)
	}
	{
		// 手动取消
		a := 1
		b := 2
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			time.Sleep(2 * time.Second)
			cancel()
		}()

		res := Add(ctx, 1, 2)
		fmt.Printf("Compute: %d+%d, result: %d\n", a, b, res)
	}
}
