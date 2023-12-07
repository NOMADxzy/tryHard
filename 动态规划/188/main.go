package main

import "fmt"

func maxProfit(k int, prices []int) int {
	m := len(prices)
	dp := make([][][2]int, k+1)
	for i := 0; i < k+1; i++ {
		dp[i] = make([][2]int, m)
	}
	dp[0][0][1] = -prices[0]
	for i := 1; i < m; i++ {
		dp[0][i][1] = max(-prices[i], dp[0][i-1][1])
	}
	ans := -1 << 31
	for curK := 1; curK <= k && curK <= m; curK++ {
		dp[curK][curK-1][1] = -1 << 31
		for i := curK; i < m; i++ {
			dp[curK][i][0] = max(dp[curK][i-1][0], dp[curK-1][i-1][1]+prices[i])
			dp[curK][i][1] = max(dp[curK][i-1][1], dp[curK][i-1][0]-prices[i])
		}
		ans = max(dp[curK][m-1][0], ans)
	}
	return ans
}

func main() {
	fmt.Println(maxProfit(2, []int{1}))
}
