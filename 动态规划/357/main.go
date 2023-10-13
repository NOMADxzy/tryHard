package main

import "fmt"

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = 9

	ans := 1 + 9
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][1] + dp[i-1][0]*(10-i)
		dp[i][1] = dp[i-1][1] * (9 - i)
		ans += dp[i][0] + dp[i][1]
	}
	return ans
}

func main() {
	fmt.Println(countNumbersWithUniqueDigits(8))
}
