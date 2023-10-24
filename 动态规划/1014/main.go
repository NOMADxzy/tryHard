package main

import "fmt"

func maxScoreSightseeingPair(values []int) int {
	m := len(values)
	dp := make([]int, m)
	dp[0] = values[0] + 0

	ans := values[0] + values[1] - 1

	for i := 1; i < m; i++ {
		if dp[i-1]+values[i]-i > ans {
			ans = dp[i-1] + values[i] - i
		}
		dp[i] = dp[i-1]
		if values[i]+i > dp[i-1] {
			dp[i] = values[i] + i
		}
	}
	return ans
}

func main() {
	fmt.Println(maxScoreSightseeingPair([]int{8, 1, 5, 2, 6}))
}
