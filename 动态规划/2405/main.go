package main

import "fmt"

func partitionString(s string) int {
	m := len(s)
	dp := make([]int, m)
	dp[0] = 1
	lastIdx := make([][2]int, 26)
	for i := 0; i < 26; i++ {
		lastIdx[i][0] = -1
		lastIdx[i][1] = -1
	}
	lastIdx[int(s[0]-'a')][1] = 0
	for i := 1; i < m; i++ {
		c := int(s[i] - 'a')
		lastIdx[c][0] = lastIdx[c][1]
		lastIdx[c][1] = i
		bestPreIdx := -1
		for j := 0; j < 26; j++ {
			bestPreIdx = max(bestPreIdx, lastIdx[j][0])
		}
		if bestPreIdx == -1 {
			dp[i] = 1
		} else {
			dp[i] = 1 + dp[bestPreIdx]
		}
	}
	return dp[m-1]
}

func main() {
	fmt.Println(partitionString("abacaba"))
}
