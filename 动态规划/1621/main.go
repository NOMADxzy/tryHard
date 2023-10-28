package main

import (
	"fmt"
)

func numberOfSets(n int, k int) int {
	MOD := 1000000007
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k+1)
		dp[i][0] = 1
	}
	sums := make([]int, n)
	for i := 0; i < n; i++ {
		sums[i] = i + 1
	}

	for curK := 1; curK <= k; curK++ {
		for i := curK; i < n; i++ {
			cnt := dp[i-1][curK]
			if curK-2 >= 0 {
				cnt = (cnt + (sums[i-1]-sums[curK-2])%MOD) % MOD
			} else {
				cnt = (cnt + sums[i-1]) % MOD
			}
			dp[i][curK] = cnt
		}
		sums[0] = dp[0][curK]
		for i := 1; i < n; i++ {
			sums[i] = sums[i-1] + dp[i][curK]
		}
	}
	return dp[n-1][k]
}

func main() {
	fmt.Println(numberOfSets(30, 7))
}
