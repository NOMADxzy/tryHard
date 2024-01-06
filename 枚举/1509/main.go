package main

import (
	"fmt"
	"sort"
)

func minDifference(nums []int) int {
	if len(nums) <= 3 {
		return 0
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	maxs := nums[len(nums)-4:]
	mins := nums[:4]
	ans := 1 << 31
	for i := 0; i < 4; i++ {
		ans = min(ans, maxs[i]-mins[i])
	}
	return ans
}

func main() {
	fmt.Println(minDifference([]int{1, 5, 0, 10, 14}))
}
