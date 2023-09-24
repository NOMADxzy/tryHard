package main

import (
	"fmt"
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	dp := make([][]int, len(nums))
	dp[0] = []int{1, -1}
	for i := 1; i < len(nums); i++ {
		dp[i] = make([]int, 2)
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	max := 0
	p := -1

	for i := 1; i < len(nums); i++ {
		maxLen, pos := 0, -1
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 && dp[j][0] > maxLen {
				maxLen = dp[j][0]
				pos = j
			}
		}
		dp[i][0] = maxLen + 1
		dp[i][1] = pos
		if dp[i][0] > max {
			max = dp[i][0]
			p = i
		}
	}
	res := make([]int, max)
	for i := 0; i < max; i++ {
		res[i] = nums[p]
		p = dp[p][1]
	}
	return res
}

func main() {
	fmt.Println(largestDivisibleSubset([]int{2, 3, 4, 9, 8}))
}
