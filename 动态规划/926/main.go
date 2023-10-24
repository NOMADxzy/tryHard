package main

import "fmt"

func minFlipsMonoIncr(s string) int {
	m := len(s)

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, 2)
	}

	if s[0] == '0' {
		dp[0][1] = 1
	} else {
		dp[0][0] = 1
	}

	for i := 1; i < m; i++ {
		best := dp[i-1][0]
		if dp[i-1][1] < best {
			best = dp[i-1][1]
		}
		if s[i] == '1' {
			dp[i][1] = best
			dp[i][0] = dp[i-1][0] + 1
		} else {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = best + 1
		}
	}
	if dp[m-1][0] < dp[m-1][1] {
		return dp[m-1][0]
	} else {
		return dp[m-1][1]
	}
}

func main() {
	fmt.Println(minFlipsMonoIncr("00011000"))
}
