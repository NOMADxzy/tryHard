package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	m := len(nums)
	dp := make([][][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int, 2)
		dp[i][0] = make([]int, 1001)
		dp[i][1] = make([]int, 1001)
	}
	if nums[0] > 0 {
		dp[0][0][nums[0]] = 1
		dp[0][1][nums[0]] = 1
	} else {
		dp[0][0][nums[0]] = 2
	}

	for i := 1; i < m; i++ {
		n := nums[i]
		for j := 0; j <= 1000; j++ {
			if dp[i-1][0][j] > 0 {
				dp[i][0][j+n] = dp[i-1][0][j]
				if j-n < 0 {
					dp[i][1][n-j] += dp[i-1][0][j]
				} else {
					dp[i][0][j-n] += dp[i-1][0][j]
				}
			}
		}
		for j := 1; j <= 1000; j++ {
			if dp[i-1][1][j] > 0 {
				dp[i][1][j+n] = dp[i-1][1][j]
				if j-n <= 0 {
					dp[i][0][n-j] += dp[i-1][1][j]
				} else {
					dp[i][1][j-n] += dp[i-1][0][j]
				}
			}
		}
	}
	ans := 0
	if target >= 0 {
		ans = dp[m-1][0][target]
	} else {
		ans = dp[m-1][1][-target]
	}
	return ans
}
func main() {
	fmt.Println(findTargetSumWays([]int{0, 1}, 1))
}
