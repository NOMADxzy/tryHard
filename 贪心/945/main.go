package main

import (
	"fmt"
	"sort"
)

func minIncrementForUnique(nums []int) int {
	ans := 0
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] <= nums[i-1] {
			ans += nums[i-1] + 1 - nums[i]
			nums[i] = nums[i-1] + 1
		}
	}
	return ans
}

func main() {
	fmt.Println(minIncrementForUnique([]int{3, 2, 1, 2, 1, 7}))
}
