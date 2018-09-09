package main

import (
	"fmt"
	"time"
)

func main() {
	var fillInternal = time.Millisecond * 10
	var caplicity = 100
	var tokenBucket = make(chan struct{}, caplicity)

	fillToken := func() {
		ticker := time.NewTicker(fillInternal)
		for {
			select {
			case <-ticker.C:
				select {
				case tokenBucket <- struct{}{}:
				default:
				}
				fmt.Println("current token cnt: ", len(tokenBucket), time.Now())
			}

		}
	}

	go fillToken()
	time.Sleep(time.Hour)
}

// func TakeAvailabe(block bool) (takenResult bool) {
// 	if block {
// 		select {
// 		case <-tokenBucket:
// 			takenResult = true
// 		}
// 	} else {
// 		select {
// 		case <-tokenBucket:
// 			takenResult = true
// 		default:
// 			takenResult = false
// 		}
// 	}

// 	return
// }
