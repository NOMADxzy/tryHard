package main

import "fmt"

func numWays(steps int, arrLen int) int {
	MOD := 1000000007
	dp1 := make([]int, 2)
	dp1[0] = 1
	maybeGet := func(dp []int, i int) int {
		if i < len(dp) {
			return dp[i]
		}
		return 0
	}
	for i := 1; i <= steps; i++ {
		lim := min(arrLen-1, min(i, steps-i))
		newDp1 := make([]int, lim+1)
		for j := 1; j <= lim; j++ {
			newDp1[j] = ((dp1[j-1]+maybeGet(dp1, j))%MOD + maybeGet(dp1, j+1)) % MOD
		}
		newDp1[0] = dp1[0]
		if len(dp1) > 1 {
			newDp1[0] = (newDp1[0] + dp1[1]) % MOD
		}
		dp1 = newDp1
	}
	return dp1[0]
}

func main() {
	fmt.Println(numWays(3, 1))
}
