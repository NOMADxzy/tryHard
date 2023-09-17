package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func rob(nums []int) int {
	m := len(nums)
	if m == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums)+1)
	dp[0] = 0
	dp[1] = 0

	for i := 2; i < len(dp); i++ {
		dp[i] = max(nums[i-1]+dp[i-2], dp[i-1])
	}
	maxProfit := dp[m]
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i < len(dp)-1; i++ {
		dp[i] = max(nums[i-1]+dp[i-2], dp[i-1])
	}
	return max(maxProfit, dp[m-1])
}

func main() {
	fmt.Println(rob([]int{1, 2}))
}
