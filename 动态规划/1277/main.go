package main

func countSquares(matrix [][]int) (ans int) {
	n := len(matrix)
	m := len(matrix[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i > 0 && j > 0 && matrix[i][j] == 1 {
				matrix[i][j] = min(matrix[i-1][j], min(matrix[i][j-1], matrix[i-1][j-1])) + 1
			}
			ans += matrix[i][j]
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
