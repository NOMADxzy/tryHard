package main

import "fmt"

func longestPalindromeSubseq(s string) int {
	m := len(s)
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}
	for l := 2; l <= m; l++ {
		for i := 0; i < m-l+1; i++ {
			left, right := i, i+l-1
			if s[left] == s[right] {
				dp[left][right] = 2
				if right-left > 1 {
					dp[left][right] += dp[left+1][right-1]
				}
			} else {
				dp[left][right] = dp[left+1][right]
				if dp[left][right-1] > dp[left][right] {
					dp[left][right] = dp[left][right-1]
				}
			}
		}
	}
	return dp[0][m-1]
}

func main() {
	fmt.Println(longestPalindromeSubseq("bbbab"))
}
