package main

import "fmt"

func minSteps(n int) int {
	dp := make([]int, n+1)
	dp[1] = 0
	for i := 2; i <= n; i++ {
		best := i
		for j := i / 2; j > 0; j-- {
			if i%j == 0 && i/j+dp[j] < best {
				best = i/j + dp[j]
			}
		}
		dp[i] = best
	}
	return dp[n]
}

func main() {
	fmt.Println(minSteps(6))
}
