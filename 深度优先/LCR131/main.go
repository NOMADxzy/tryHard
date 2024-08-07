package main

import "fmt"

func cuttingBamboo(bamboo_len int) int {
	dp := map[int]int{1: 1}
	var dfs func(l int) int
	dfs = func(l int) int {
		if v, ok := dp[l]; ok {
			return v
		}
		maxVal := l
		for mid := 0; mid+1 <= l-mid-1; mid++ {
			maxVal = max(maxVal, dfs(mid+1)*dfs(l-mid-1))
		}
		dp[l] = maxVal
		return maxVal
	}
	ans := 0
	for i := 0; i+1 <= bamboo_len-1-i; i++ {
		ans = max(ans, dfs(i+1)*dfs(bamboo_len-i-1))
	}
	return ans
}

func main() {
	fmt.Println(cuttingBamboo(12))
}
