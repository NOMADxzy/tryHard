package main

import "fmt"

func minCut(s string) int {
	m := len(s)
	dp := make([][]bool, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]bool, m)
		dp[i][i] = true
	}
	for l := 1; l < m; l++ {
		for i := 0; i < m-l; i++ {
			if s[i] == s[i+l] && (l == 1 || dp[i+1][i+l-1]) {
				dp[i][i+l] = true
			}
		}
	}
	leftCuts := make([]int, m)
	leftCuts[0] = 0
	for i := 1; i < m; i++ {
		if !dp[0][i] {
			best := i
			for j := i; j > 0; j-- {
				if dp[j][i] && leftCuts[j-1]+1 < best {
					best = leftCuts[j-1] + 1
				}
			}
			leftCuts[i] = best
		}
	}
	return leftCuts[m-1]
}

func main() {
	fmt.Println(minCut("aab"))
}
