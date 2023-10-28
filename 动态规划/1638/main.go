package main

import "fmt"

func countSubstrings(s string, t string) int {
	m, n := len(s), len(t)
	dp := make([][][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([][]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	ans := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] != t[j-1] {
				dp[i][j][1] = dp[i-1][j-1][0] + 1
			} else {
				dp[i][j][0] = dp[i-1][j-1][0] + 1

				dp[i][j][1] = dp[i-1][j-1][1]
				//if dp[i-1][j][1] > dp[i][j][1] {
				//	dp[i][j][1] = dp[i-1][j][1]
				//}
			}
			ans += dp[i][j][1]
		}
	}
	return ans
}

func main() {
	fmt.Println(countSubstrings("aba", "baba"))
}
