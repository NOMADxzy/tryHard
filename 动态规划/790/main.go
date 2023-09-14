package main

import "fmt"

func numTilings(n int) int {

	if n <= 1 {
		return 1
	}
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0], dp[0][1] = 0, 1
	dp[1][0], dp[1][1] = 0, 1

	for i := 2; i <= n; i++ {
		dp[i][0] = (dp[i-1][0] + dp[i-2][1]*2) % 1000000007
		dp[i][1] = (dp[i-1][0] + dp[i-2][1] + dp[i-1][1]) % 1000000007
	}
	return dp[n][1]
}

func main() {
	fmt.Println(numTilings(30))
}
