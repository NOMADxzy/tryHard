package main

import "fmt"

func maxProfit(prices []int) int {
	dp := make([]int, len(prices))
	dp[0] = prices[0]

	max := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < dp[i-1] {
			dp[i] = prices[i]
		} else {
			dp[i] = dp[i-1]
			if max < prices[i]-dp[i-1] {
				max = prices[i] - dp[i-1]
			}
		}
	}
	return max
}

func main() {
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
