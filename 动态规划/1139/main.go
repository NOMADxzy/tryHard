package main

import "fmt"

func largest1BorderedSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	dp := make([][][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, 4)
		}
	}
	for i := 0; i < m; i++ {
		if grid[i][0] == 1 {
			dp[i][0][3] = 1
		}
		if grid[i][n-1] == 1 {
			dp[i][n-1][1] = 1
		}
	}
	for j := 0; j < n; j++ {
		if grid[0][j] == 1 {
			dp[0][j][0] = 1
		}
		if grid[m-1][j] == 1 {
			dp[m-1][j][2] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dp[i][j][0] = dp[i-1][j][0] + 1
			}
			if grid[m-1-i][j] == 1 {
				dp[m-1-i][j][2] = dp[m-i][j][2] + 1
			}
		}
	}
	for j := 1; j < n; j++ {
		for i := 0; i < m; i++ {
			if grid[i][j] == 1 {
				dp[i][j][3] = dp[i][j-1][3] + 1
			}
			if grid[i][n-1-j] == 1 {
				dp[i][n-1-j][1] = dp[i][n-j][1] + 1
			}
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for rightBottomX, rightBottomY := i+ans, j+ans; ; {
				if rightBottomX >= m || rightBottomY >= n {
					break
				} else if dp[i][j][1] > rightBottomX-i && dp[i][j][2] > rightBottomX-i && dp[rightBottomX][rightBottomY][0] > rightBottomX-i && dp[rightBottomX][rightBottomY][3] > rightBottomX-i {
					ans = rightBottomX - i + 1
				}
				rightBottomX++
				rightBottomY++
			}
		}
	}
	return ans * ans
}

func main() {
	fmt.Println(largest1BorderedSquare([][]int{{0}}))
}
