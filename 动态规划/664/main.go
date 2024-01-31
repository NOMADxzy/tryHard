package main

import "fmt"

func strangePrinter(s string) int {
	m := len(s)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
		dp[i][i] = 1
	}
	for l := 1; l < m; l++ {
		for i := 0; i < m-l; i++ {
			best := m
			if s[i] == s[i+l] {
				best = dp[i][i+l-1]
			} else {
				for mid := i; mid < i+l; mid++ {
					best = min(best, dp[i][mid]+dp[mid+1][i+l])
				}
			}
			dp[i][i+l] = best
		}
	}
	return dp[0][m-1]
}

func main() {
	fmt.Println(strangePrinter("aba"))
}
