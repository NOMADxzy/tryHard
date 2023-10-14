package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			steps := 0
			if word1[i-1] != word2[j-1] {
				steps++
				if dp[i-1][j] > dp[i][j-1] {
					steps += dp[i][j-1]
				} else {
					steps += dp[i-1][j]
				}
			} else {
				steps = dp[i-1][j-1]
			}
			dp[i][j] = steps
		}
	}
	return dp[m][n]
}

func main() {
	fmt.Println(minDistance("eat", "sea"))
}
