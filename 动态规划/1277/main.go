package main

import "fmt"

func countSquares(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	cnt := 0
	for size := 1; ; size++ {

		curCnt := 0
		for i := size - 1; i < m; i++ {
			for j := size - 1; j < n; j++ {
				if matrix[i][j] == 1 && (size == 1 || dp[i-1][j-1] >= size-1 && dp[i-1][j] >= size-1 && dp[i][j-1] >= size-1) {
					dp[i][j] = size
					curCnt++
				}
			}
		}
		cnt += curCnt
		if curCnt == 0 {
			break
		}
	}
	return cnt
}

func main() {
	fmt.Println(countSquares([][]int{{0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}}))
}
