package main

import "fmt"

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	m := len(nums)
	dp := make([][2]int, m+1)
	sum := 0
	for i := 1; i <= m; i++ {
		if nums[i-1] > right {
			continue
		} else if nums[i-1] < left {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = dp[i-1][1] + 1
		} else {
			dp[i][0] = dp[i-1][1] + 1
			dp[i][1] = dp[i-1][1] + 1
		}
		sum += dp[i][0]
	}
	return sum
}

func main() {
	fmt.Println(numSubarrayBoundedMax([]int{73, 55, 36, 5, 55, 14, 9, 7, 72, 52}, 32, 69))
}
