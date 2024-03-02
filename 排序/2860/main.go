package main

import (
	"fmt"
	"sort"
)

func countWays(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ans := 0
	for i := 0; i <= len(nums); i++ {
		if i < len(nums) && nums[i] <= i {
			continue
		}
		if i > 0 && nums[i-1] >= i {
			continue
		}
		ans++
	}
	return ans
}

func main() {
	fmt.Println(countWays([]int{6, 0, 3, 3, 6, 7, 2, 7}))
}
