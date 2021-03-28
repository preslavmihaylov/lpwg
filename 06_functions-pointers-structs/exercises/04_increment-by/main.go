package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	incrementBy(nums, 5)
	fmt.Println(nums)
}

func incrementBy(nums []int, n int) {
	for i := range nums {
		nums[i] += n
	}
}
