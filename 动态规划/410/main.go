package main

import "fmt"

func splitArray(nums []int, k int) int {
	m := len(nums)
	dp := make([]int, m)
	dp[0] = nums[0]
	for i := 1; i < m; i++ {
		dp[i] = dp[i-1] + nums[i]
	}
	for j := 1; j < k; j++ {
		newDp := make([]int, m)
		for i := j; i < m; i++ {
			best := 1 << 31
			tmp := 0
			for t := i; t >= j; t-- {
				tmp += nums[t]
				if v := max(tmp, dp[t-1]); v <= best { //TODO 易错，不能写 v < best, 1000,100,100,100,1 这种情况前几次v都没变，但最优值在最前面，不能提前break，
					best = v
				} else {
					break
				}
			}
			newDp[i] = best
		}
		dp = newDp
	}
	return dp[m-1]
}

func main() {
	fmt.Println(splitArray([]int{7, 2, 5, 10, 8}, 2))
}
