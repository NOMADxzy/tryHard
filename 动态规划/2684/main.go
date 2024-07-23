package main

import "fmt"

func maxMoves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, m)
	ans := 0
	for j := 1; j < n; j++ {
		newDp := make([]int, m)
		for i := 0; i < m; i++ {
			if i-1 >= 0 && grid[i][j] > grid[i-1][j-1] && (dp[i-1] > 0 || j == 1) {
				newDp[i] = max(newDp[i], dp[i-1]+1)
			}
			if grid[i][j] > grid[i][j-1] && (dp[i] > 0 || j == 1) {
				newDp[i] = max(newDp[i], dp[i]+1)
			}
			if i+1 < m && grid[i][j] > grid[i+1][j-1] && (dp[i+1] > 0 || j == 1) {
				newDp[i] = max(newDp[i], dp[i+1]+1)
			}
			ans = max(ans, newDp[i])
		}
		dp = newDp
	}
	return ans
}

func main() {
	fmt.Println(maxMoves([][]int{
		{2, 4, 3, 5},
		{5, 4, 9, 3},
		{3, 4, 2, 11},
		{10, 9, 13, 15},
	}))
}
