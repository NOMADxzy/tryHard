package main

import "fmt"

func maxProfit(prices []int) int {
	dp := make([]int, len(prices))
	dp[0] = prices[0]

	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[i-1] {
			profit += prices[i-1] - dp[i-1]
			dp[i] = prices[i]
		} else {
			dp[i] = dp[i-1]
			if i == len(prices)-1 {
				profit += prices[i] - dp[i-1]
			}
		}
	}
	return profit
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
}
