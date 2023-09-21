package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[1][0] = max(0, prices[1]-prices[0])
	dp[1][1] = max(-prices[0], -prices[1])

	for i := 2; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-2][0]-prices[i], dp[i-1][1])
	}
	return dp[len(prices)-1][0]
}
func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
}
