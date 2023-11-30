package main

import "fmt"

func numDistinct(s string, t string) int {
	MOD := 1000000007
	m, n := len(s), len(t)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i < m+1; i++ {
		dp[0][i] = 1
	}

	for j := 1; j <= n; j++ {
		for i := j; i <= m; i++ {
			dp[j][i] = dp[j][i-1]
			if s[i-1] == t[j-1] {
				dp[j][i] = (dp[j-1][i-1] + dp[j][i]) % MOD
			}
		}
	}
	return dp[n][m]
}

func main() {
	fmt.Println(numDistinct("babgbag", "bagbabgbag"))
}
