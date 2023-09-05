package main

import "fmt"

func numSquares(n int) int {
	dp := make([]int, n+1)

	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		min := dp[i-1] + 1
		for j := 2; j*j <= i; j++ {
			step := dp[i-j*j] + 1
			if step < min {
				min = step
			}
		}
		dp[i] = min
	}
	return dp[n]
}

func main() {
	fmt.Println(numSquares(1))
}
