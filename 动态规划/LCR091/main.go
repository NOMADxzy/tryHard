package main

import "fmt"

func minCost(costs [][]int) int {
	m := len(costs)
	dp := make([][3]int, m)
	dp[0][0] = costs[0][0]
	dp[0][1] = costs[0][1]
	dp[0][2] = costs[0][2]
	for i := 1; i < m; i++ {
		dp[i][0] = min(dp[i-1][1]+costs[i][0], dp[i-1][2]+costs[i][0])
		dp[i][1] = min(dp[i-1][0]+costs[i][1], dp[i-1][2]+costs[i][1])
		dp[i][2] = min(dp[i-1][0]+costs[i][2], dp[i-1][1]+costs[i][2])
	}
	return min(dp[m-1][0], min(dp[m-1][1], dp[m-1][2]))
}

func main() {
	fmt.Println(minCost([][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}))
}
