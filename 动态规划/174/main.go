package main

import "fmt"

func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[m-1][n-1] = max(-dungeon[m-1][n-1]+1, 0)
	for i := m - 2; i >= 0; i-- {
		dp[i][n-1] = max(max(-dungeon[i][n-1]+1, 0), dp[i+1][n-1]-dungeon[i][n-1])
	}
	for j := n - 2; j >= 0; j-- {
		dp[m-1][j] = max(max(-dungeon[m-1][j]+1, 0), dp[m-1][j+1]-dungeon[m-1][j])
	}
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			val1 := max(max(-dungeon[i][j]+1, 0), dp[i+1][j]-dungeon[i][j])
			val2 := max(max(-dungeon[i][j]+1, 0), dp[i][j+1]-dungeon[i][j])
			dp[i][j] = min(val1, val2)
		}
	}
	return max(dp[0][0], 1)
}

func main() {
	mat := [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}
	fmt.Println(calculateMinimumHP(mat))
}
