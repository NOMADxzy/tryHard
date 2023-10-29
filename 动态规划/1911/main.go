package main

import "fmt"

func maxAlternatingSum(nums []int) int64 {
	m := len(nums)
	dp := make([]int64, 2)
	nextDp := make([]int64, 2)
	dp[0] = 0
	dp[1] = int64(nums[0])
	copy(nextDp, dp)

	for i := 1; i < m; i++ {

		if dp[1]-int64(nums[i]) > nextDp[0] {
			nextDp[0] = dp[1] - int64(nums[i])
		}
		if dp[0]+int64(nums[i]) > nextDp[1] {
			nextDp[1] = dp[0] + int64(nums[i])
		}
		copy(dp, nextDp)
	}
	if dp[0] > dp[1] {
		return dp[0]
	} else {
		return dp[1]
	}
}

func main() {
	fmt.Println(maxAlternatingSum([]int{4, 2, 5, 3}))
}
