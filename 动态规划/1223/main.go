package main

import "fmt"

func dieSimulator(n int, rollMax []int) int {
	MOD := 1000000007
	dp := make([][]int, 6)
	for i := 0; i < 6; i++ {
		dp[i] = make([]int, rollMax[i])
		dp[i][0] = 1
	}
	for i := 1; i < n; i++ {
		newDp := make([][]int, 6)
		tmp := make([]int, 6)
		for j := 0; j < 6; j++ {
			newDp[j] = make([]int, rollMax[j])
			sum := 0
			for k := 1; k < rollMax[j] && dp[j][k-1] > 0; k++ {
				newDp[j][k] = dp[j][k-1]
				//tmp = (tmp + dp[j][k-1]) % MOD
				sum = (sum + dp[j][k-1]) % MOD
			}
			sum = (sum + dp[j][rollMax[j]-1]) % MOD
			for p := 0; p < 6; p++ {
				if p != j {
					tmp[p] = (tmp[p] + sum) % MOD
				}
			}
		}
		for j := 0; j < 6; j++ {
			newDp[j][0] = tmp[j]
		}
		dp = newDp
	}
	ans := 0
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			ans = (ans + dp[i][j]) % MOD
		}
	}
	return ans
}

func main() {
	fmt.Println(dieSimulator(20, []int{8, 5, 10, 8, 7, 2}))
}
