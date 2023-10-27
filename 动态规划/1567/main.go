package main

import "fmt"

func getMaxLen(nums []int) int {
	m := len(nums)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, 2)
	}

	ans := 0
	if nums[0] > 0 {
		dp[0][0] = 1
		ans = 1
	} else if nums[0] < 0 {
		dp[0][1] = 1
	}

	for i := 1; i < m; i++ {
		if nums[i] == 0 {
			continue
		} else if nums[i] > 0 {
			dp[i][0] = dp[i-1][0] + 1
			if dp[i-1][1] > 0 {
				dp[i][1] = dp[i-1][1] + 1
			} else {
				dp[i][1] = 0
			}
		} else {
			dp[i][1] = dp[i-1][0] + 1
			if dp[i-1][1] > 0 {
				dp[i][0] = dp[i-1][1] + 1
			} else {
				dp[i][0] = 0
			}
		}
		if dp[i][0] > ans {
			ans = dp[i][0]
		}
	}
	return ans
}

func main() {
	fmt.Println(getMaxLen([]int{0, 1, -2, -3, -4}))
}
