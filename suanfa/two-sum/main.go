package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var res []int
	for at := 0; at < len(nums); at++ {
		for bt := at + 1; bt < len(nums); bt++ {
			if nums[at]+nums[bt] == target {
				res = append(res, at, bt)
				break
			}
		}
	}

	return res
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
