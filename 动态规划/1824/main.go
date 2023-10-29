package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minSideJumps(obstacles []int) int {
	n := len(obstacles) - 1
	MAXINT := 10000000000

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 3)
	}
	dp[0][0], dp[0][2] = 1, 1

	for i := 1; i <= n; i++ {
		for side := 0; side < 3; side++ {
			if obstacles[i]-1 == side || obstacles[i-1]-1 == side {
				dp[i][side] = MAXINT
			} else {
				dp[i][side] = dp[i-1][side]
			}
		}
		dp[i][0] = min(dp[i][0], min(dp[i][2], dp[i][1])+1)
		dp[i][1] = min(dp[i][1], min(dp[i][2], dp[i][0])+1)
		dp[i][2] = min(dp[i][2], min(dp[i][0], dp[i][1])+1)
		if obstacles[i]-1 >= 0 {
			dp[i][obstacles[i]-1] = MAXINT
		}
	}
	return min(dp[n][0], min(dp[n][1], dp[n][2]))
}

func main() {
	fmt.Println(minSideJumps([]int{0, 1, 2, 3, 0}))
}
