package main

import "fmt"

func maxProductPath(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, 2)
			dp[i][j][0] = -1
		}
	}
	if grid[m-1][n-1] >= 0 {
		dp[m-1][n-1][0] = grid[m-1][n-1]
	} else {
		dp[m-1][n-1][0] = -1
		dp[m-1][n-1][1] = grid[m-1][n-1]
	}

	exist0 := false

	MOD := 1000000007
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == 0 {
				exist0 = true
			}
			maxPos, maxNeg := -1, 0
			if j < n-1 {
				if dp[i][j+1][0] >= 0 {
					val := (grid[i][j] * dp[i][j+1][0])
					if val > maxPos {
						maxPos = val
					} else if val < maxNeg {
						maxNeg = val
					}
				}
				if dp[i][j+1][1] < 0 {
					val := (grid[i][j] * dp[i][j+1][1])
					if val > maxPos {
						maxPos = val
					} else if val < maxNeg {
						maxNeg = val
					}
				}
			}
			if i < m-1 {
				if dp[i+1][j][0] >= 0 {
					val := (grid[i][j] * dp[i+1][j][0])
					if val > maxPos {
						maxPos = val
					} else if val < maxNeg {
						maxNeg = val
					}
				}
				if dp[i+1][j][1] < 0 {
					val := (grid[i][j] * dp[i+1][j][1])
					if val > maxPos {
						maxPos = val
					} else if val < maxNeg {
						maxNeg = val
					}
				}
			}
			if maxPos >= 0 {
				dp[i][j][0] = maxPos
			}
			if maxNeg < 0 {
				dp[i][j][1] = maxNeg
			}
		}
	}
	if dp[0][0][0] < 0 && exist0 {
		return 0
	}
	return dp[0][0][0] % MOD
}

func main() {
	grid := [][]int{{-1, 3, 0}, {3, -2, 3}, {-1, 1, -4}}
	fmt.Println(maxProductPath(grid))
}
