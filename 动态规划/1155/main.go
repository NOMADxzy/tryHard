package main

import "fmt"

func numRollsToTarget(n int, k int, target int) int {
	MOD := 1000000007
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, target+1)
	}

	for i := 1; i <= k && i <= target; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < n; i++ {
		for j := i + 1; j <= target; j++ {
			ways := 0
			for point := 1; point <= k && j-point >= 0; point++ {
				ways = (ways + dp[i-1][j-point]) % MOD
			}
			dp[i][j] = ways
		}
	}

	return dp[n-1][target]
}

func main() {
	fmt.Println(numRollsToTarget(30, 30, 500))
}
