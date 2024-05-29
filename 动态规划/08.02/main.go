package main

import "fmt"

func pathWithObstacles(obstacleGrid [][]int) [][]int {
	if obstacleGrid[0][0] == 1 {
		return make([][]int, 0)
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]bool, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]bool, n)
	}
	dp[0][0] = true
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				if i-1 >= 0 && dp[i-1][j] || j-1 >= 0 && dp[i][j-1] {
					dp[i][j] = true
				}
			}
		}
	}
	if !dp[m-1][n-1] {
		return make([][]int, 0)
	}
	var ans [][]int
	x, y := m-1, n-1
	ans = append(ans, []int{m - 1, n - 1})
	for x != 0 || y != 0 {
		if x-1 >= 0 && dp[x-1][y] {
			x -= 1
			ans = append(ans, []int{x, y})
		} else if y-1 >= 0 && dp[x][y-1] {
			y -= 1
			ans = append(ans, []int{x, y})
		}
	}
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return ans
}

func main() {
	grid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	fmt.Println(pathWithObstacles(grid))
}
