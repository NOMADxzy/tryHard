package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	dp := []int{triangle[0][0]}
	m := len(triangle)
	for i := 1; i < m; i++ {
		newDp := make([]int, i+1)
		newDp[0] = dp[0] + triangle[i][0]
		for j := 1; j < len(newDp)-1; j++ {
			newDp[j] = min(dp[j-1], dp[j]) + triangle[i][j]
		}
		newDp[len(newDp)-1] = dp[len(dp)-1] + triangle[i][len(triangle[i])-1]
		dp = newDp
	}
	ans := dp[0]
	for i := 0; i < len(dp); i++ {
		ans = min(ans, dp[i])
	}
	return ans
}

func main() {
	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
}
