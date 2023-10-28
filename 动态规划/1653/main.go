package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minimumDeletions(s string) int {
	m := len(s)
	dp := make([]int, 2)

	if s[0] == 'a' {
		dp[1] = -1
	} else {
		dp[0] = 1
	}
	for i := 1; i < m; i++ {
		nextDp := make([]int, 2)
		if dp[1] < 0 {
			if s[i] == 'a' {
				continue
			} else {
				nextDp[0] = dp[0] + 1
				nextDp[1] = dp[0]
			}
		} else {
			minA, minB := m, m
			if s[i] == 'a' {
				minA = min(minA, dp[0])
				minB = dp[1] + 1
			} else {
				minA = min(minA, dp[0]+1)
				minB = min(dp[0], dp[1])
			}
			nextDp[0] = minA
			nextDp[1] = minB
		}
		dp = nextDp
	}
	if dp[1] == -1 {
		return 0
	}
	return min(dp[0], dp[1])
}

func main() {
	fmt.Println(minimumDeletions("aababbab"))
}
