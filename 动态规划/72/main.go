package main

import "fmt"

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	dp[0][0] = 0
	for i := 1; i < len(word1)+1; i++ {
		dp[i][0] = i
	}
	for i := 1; i < len(word2)+1; i++ {
		dp[0][i] = i
	}

	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				choice1 := 1 + dp[i-1][j-1]
				choice2 := 1 + dp[i][j-1]
				choice3 := 1 + dp[i-1][j]

				dp[i][j] = min(min(choice1, choice2), choice3)
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func main() {
	fmt.Println(minDistance("intention", "execution"))
}
