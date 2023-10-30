package main

import "fmt"

func minimizeTheDifference(mat [][]int, target int) int {
	m, n := len(mat), len(mat[0])
	rowMin := make([]int, m)
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			if rowMin[i] == 0 || mat[i][j] < rowMin[i] {
				rowMin[i] = mat[i][j]
			}
		}
	}
	minBottomSum := make([]int, m+1)
	for i := m - 1; i >= 0; i-- {
		minBottomSum[i] = minBottomSum[i+1] + rowMin[i]
	}

	dp := make([]bool, target+1)
	ans := 1 << 31

	dp[0] = true

	for r := 0; r < m; r++ {
		nextDp := make([]bool, target+1)
		for t := 0; t <= target; t++ {
			if dp[t] {
				for j := 0; j < n; j++ {
					if t+mat[r][j] >= target {
						if val := t + mat[r][j] + minBottomSum[r+1] - target; val < ans {
							ans = val
						}
					} else {
						nextDp[t+mat[r][j]] = true
					}
				}
			}
		}
		dp = nextDp
	}
	for i := target; i > 0; i-- {
		if dp[i] {
			if target-i < ans {
				ans = target - i
			}
			break
		}
	}
	return ans
}

func main() {
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(minimizeTheDifference(mat, 13))
}
