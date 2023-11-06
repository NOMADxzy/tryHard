package main

import (
	"fmt"
	"sort"
)

func largestSubmatrix(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n)
	}
	for i := m - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				dp[i][j] = dp[i+1][j] + 1
			}
		}
	}
	ans := 0
	heights := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			heights[j] = dp[i][j]
		}
		sort.Slice(heights, func(i, j int) bool {
			return heights[i] > heights[j]
		})
		for w := 1; w <= n && heights[w-1] > 0; w++ {
			if val := w * heights[w-1]; val > ans {
				ans = val
			}
		}
	}
	return ans
}
func main() {
	matrix := [][]int{{0, 0, 1}, {1, 1, 1}, {1, 0, 1}}
	fmt.Println(largestSubmatrix(matrix))
}
