package main

import "fmt"

func numberOfStableArrays(zero int, one int, limit int) int {
	dp := make([][][2]int, zero+1)
	mod := int(1e9 + 7)
	for i := 0; i <= zero; i++ {
		dp[i] = make([][2]int, one+1)
	}
	for i := 0; i <= min(zero, limit); i++ {
		dp[i][0][0] = 1
	}
	for j := 0; j <= min(one, limit); j++ {
		dp[0][j][1] = 1
	}
	for i := 1; i <= zero; i++ {
		for j := 1; j <= one; j++ {
			if i > limit {
				dp[i][j][0] = dp[i-1][j][0] + dp[i-1][j][1] - dp[i-limit-1][j][1]
			} else {
				dp[i][j][0] = dp[i-1][j][0] + dp[i-1][j][1]
			}
			dp[i][j][0] = (dp[i][j][0]%mod + mod) % mod
			if j > limit {
				dp[i][j][1] = dp[i][j-1][1] + dp[i][j-1][0] - dp[i][j-limit-1][0]
			} else {
				dp[i][j][1] = dp[i][j-1][1] + dp[i][j-1][0]
			}
			dp[i][j][1] = (dp[i][j][1]%mod + mod) % mod
		}
	}
	return (dp[zero][one][0] + dp[zero][one][1]) % mod
}
func numberOfStableArrays1(zero int, one int, limit int) int {
	dp := map[int]int{}
	MOD := 1000000007
	var dfs func(acc0, acc1, remain0, remain1 int) int
	dfs = func(acc0, acc1, remain0, remain1 int) int {
		key := 201*(201*(201*acc0+acc1)+remain0) + remain1
		if v, ok := dp[key]; ok {
			return v
		}
		if remain0 == 0 && remain1 == 0 {
			return 1
		}
		cnt := 0
		if remain0 > 0 && acc0 < limit {
			cnt = (dfs(acc0+1, 0, remain0-1, remain1) + cnt) % MOD
		}
		if remain1 > 0 && acc1 < limit {
			cnt = (dfs(0, acc1+1, remain0, remain1-1) + cnt) % MOD
		}
		dp[key] = cnt
		return cnt
	}
	return dfs(0, 0, zero, one)
}

func main() {
	fmt.Println(numberOfStableArrays(200, 200, 100))
}
