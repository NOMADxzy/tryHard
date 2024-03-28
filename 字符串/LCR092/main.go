package main

import "fmt"

func minFlipsMonoIncr(s string) int {
	dp := make([]int, 2)
	if s[0] == '0' {
		dp[1] = 1
	} else {
		dp[0] = 1
	}
	for i := 1; i < len(s); i++ {
		newDp := make([]int, 2)
		to0, to1 := 0, 0
		if s[i] == '0' {
			to1++
		} else {
			to0++
		}
		newDp[0] = dp[0] + to0
		newDp[1] = min(dp[0], dp[1]) + to1
		dp = newDp
	}
	return min(dp[0], dp[1])
}

func main() {
	fmt.Println(minFlipsMonoIncr("010110"))
}
