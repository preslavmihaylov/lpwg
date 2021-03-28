package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(sublist(nums, 1, 12))
}

func sublist(nums []int, start, end int) []int {
	if start >= len(nums) {
		return []int{}
	} else if end >= len(nums) {
		end = len(nums)
	}

	return nums[start:end]
}
