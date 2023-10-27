package main

import "fmt"

func numSubmat(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	dp := make([][][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, m)
		}
	}

	ans := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				for h := 1; h <= i+1 && mat[i+1-h][j] == 1; h++ {
					if j == 0 {
						dp[i][j][h-1] = 1
					} else {
						dp[i][j][h-1] = dp[i][j-1][h-1] + 1
					}
					ans += dp[i][j][h-1]
				}
			}
		}
	}
	return ans
}

func main() {
	mat := [][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}
	fmt.Println(numSubmat(mat))
}
