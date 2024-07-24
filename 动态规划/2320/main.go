package main

import "fmt"

func countHousePlacements(n int) int {
	dp := make([][2]int, 2)
	dp[0][0] = 1
	dp[0][1] = 1
	dp[1][0] = 1
	dp[1][1] = 1

	ans := 0
	MOD := 1000000007
	for i := 1; i < n; i++ {
		newDp := make([][2]int, 2)
		newDp[0][0] = (dp[0][0] + dp[0][1] + dp[1][0] + dp[1][1]) % MOD
		newDp[0][1] = (dp[1][0] + dp[0][0]) % MOD
		newDp[1][0] = (dp[0][1] + dp[0][0]) % MOD
		newDp[1][1] = dp[0][0]
		dp = newDp
	}
	ans = (dp[0][0] + (dp[0][1]+(dp[1][0]+dp[1][1])%MOD)%MOD) % MOD
	return ans
}

func main() {
	fmt.Println(countHousePlacements(2))
}
