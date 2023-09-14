package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxProfit(prices []int, fee int) int {
	dp := make([][]int, len(prices))

	dp[0] = []int{0, -prices[0]}
	for i := 1; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}

	for i := 1; i < len(dp); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	return dp[len(dp)-1][0]
}

func main() {
	fmt.Println(maxProfit([]int{1, 3, 7, 5, 10, 3}, 3))
}
