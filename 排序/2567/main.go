package main

import (
	"fmt"
	"sort"
)

func minimizeSum(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	if len(nums) <= 3 {
		return 0
	}
	x1 := nums[len(nums)-1] - nums[2]
	x2 := nums[len(nums)-3] - nums[0]
	x3 := nums[len(nums)-2] - nums[1]

	return min(x1, min(x2, x3))
}

func main() {
	fmt.Println(minimizeSum([]int{31, 25, 72, 79, 74, 65}))
}
