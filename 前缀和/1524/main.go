package main

import "fmt"

func numOfSubarrays(arr []int) int {
	MOD_NUM := 1000000007
	m := len(arr)
	dp := make([][]int, 2)
	dp[0] = make([]int, 2)
	dp[1] = make([]int, 2)
	ans := 0
	for i := 0; i < m; i++ {
		if arr[i]%2 == 0 {
			dp[1][0] = (dp[0][0] + 1) % MOD_NUM
			dp[1][1] = dp[0][1]
		} else {
			dp[1][1] = (dp[0][0] + 1) % MOD_NUM
			dp[1][0] = dp[0][1]
		}
		ans = (ans + dp[1][1]%MOD_NUM) % MOD_NUM
		copy(dp[0], dp[1])
	}
	return ans
}

func main() {
	fmt.Println(numOfSubarrays([]int{1, 3, 5}))
}
