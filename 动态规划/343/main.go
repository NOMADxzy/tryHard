package main

import "fmt"

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1

	for i := 2; i <= n; i++ {
		max := i
		if i == n {
			max = 1
		}
		for j := i - 1; j > 0; j-- {
			if dp[j]*(i-j) > max {
				max = dp[j] * (i - j)
			}
		}
		dp[i] = max
	}
	return dp[n]
}

func main() {
	fmt.Println(integerBreak(3))
}
