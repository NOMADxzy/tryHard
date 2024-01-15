package main

import "fmt"

// 动态规划
func maxSumMinProduct1(nums []int) int {
	n := len(nums)
	best := int64(0)
	dp := make([][][2]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int64, n)
		dp[i][i][0] = int64(nums[i])
		dp[i][i][1] = int64(nums[i])
		best = max(best, dp[i][i][0]*dp[i][i][1]) //TODO 容易漏掉
	}

	for l := 1; l < n; l++ {
		for i := 0; i < n-l; i++ {
			minVal, sumVal := min(dp[i][i][0], dp[i+1][i+l][0]), dp[i][i][1]+dp[i+1][i+l][1]
			dp[i][i+l][0], dp[i][i+l][1] = minVal, sumVal
			best = max(best, minVal*sumVal)
		}
	}
	MOD := int64(1000000007)
	return int(best % MOD)
}

func maxSumMinProduct(nums []int) int {
	n := len(nums)
	best := int64(0)
	dp := make([][][2]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int64, n)
		dp[i][i][0] = int64(nums[i])
		dp[i][i][1] = int64(nums[i])
		best = max(dp[i][i][0], dp[i][i][1])
	}

	for l := 1; l < n; l++ {
		for i := 0; i < n-l; i++ {
			minVal, sumVal := min(dp[i][i][0], dp[i+1][i+l][0]), dp[i][i][1]+dp[i+1][i+l][1]
			dp[i][i+l][0], dp[i][i+l][1] = minVal, sumVal
			best = max(best, minVal*sumVal)
		}
	}
	MOD := int64(1000000007)
	return int(best % MOD)
}
func main() {
	fmt.Println(maxSumMinProduct([]int{1, 2, 3, 2}))
}
