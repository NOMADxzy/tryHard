package main

import (
	"fmt"
	"sort"
)

func combinationSum4(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	if nums[0] > target {
		return 0
	}
	dp := make([]int, target+1)
	dp[0] = 1
	for i := nums[0]; i <= target; i++ {
		curAns := 0
		for j := 0; j < len(nums) && nums[j] <= i; j++ {
			curAns += dp[i-nums[j]]
		}
		dp[i] = curAns
	}
	return dp[target]
}

func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
}
