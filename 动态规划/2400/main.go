package main

import "fmt"

// endPos-k, endPos, endPos+k

func numberOfWays(startPos int, endPos int, k int) int {
	dp := make([]int, 2*k+1)
	base := endPos - k
	MOD := 1000000007
	if startPos < endPos-k || startPos > endPos+k {
		return 0
	}
	dp[startPos-base] = 1
	for t := k - 1; t >= 0; t-- {
		newDp := make([]int, 2*t+1)
		for i := 0; i < len(newDp); i++ {
			newDp[i] = (dp[i] + dp[i+2]) % MOD
		}
		dp = newDp
	}
	return dp[0]
}

func main() {
	fmt.Println(numberOfWays(1, 2, 3))
}
