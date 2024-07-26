package main

func maxScore(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	ans := grid[0][1] - grid[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 {
				v := dp[i-1][j] + (grid[i][j] - grid[i-1][j])
				ans = max(ans, v)
				if v > dp[i][j] {
					dp[i][j] = v
				}
			}
			if j > 0 {
				v := dp[i][j-1] + (grid[i][j] - grid[i][j-1])
				ans = max(ans, v)
				if v > dp[i][j] {
					dp[i][j] = v
				}
			}
		}
	}
	return ans
}
