package main

import "fmt"

func main() {
	var nums = []int{2, 2, 1}
	sn := singleNumber(nums)
	fmt.Println("appear one time: ", sn)
}

func singleNumber(nums []int) int {
	var ret = 0
	for _, v := range nums {
		ret ^= v
	}

	return ret
}
