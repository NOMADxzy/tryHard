package main

import "fmt"

// aeiou
func countVowelPermutation(n int) int {
	MOD := 1000000007
	dp := make([]int, 5)
	for i := 0; i < 5; i++ {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		newDp := make([]int, 5)
		newDp[0] = ((dp[1]+dp[2])%MOD + dp[4]) % MOD
		newDp[1] = (dp[0] + dp[2]) % MOD
		newDp[2] = (dp[1] + dp[3]) % MOD
		newDp[3] = dp[2]
		newDp[4] = (dp[2] + dp[3]) % MOD
		dp = newDp
	}
	ans := 0
	for i := 0; i < len(dp); i++ {
		ans = (ans + dp[i]) % MOD
	}
	return ans
}

func main() {
	fmt.Println(countVowelPermutation(5))
}
