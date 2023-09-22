package main

import "fmt"

func choose(arr []int, j int) int {
	min := arr[j]
	if j-1 >= 0 && arr[j-1] < min {
		min = arr[j-1]
	}
	if j+1 < len(arr) && arr[j+1] < min {
		min = arr[j+1]
	}
	return min
}

func minFallingPathSum(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		dp[0][j] = matrix[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[1][j] = choose(dp[0], j) + matrix[i][j]
		}
		copy(dp[0], dp[1])
	}
	best := dp[0][0]
	for i := 1; i < n; i++ {
		if dp[0][i] < best {
			best = dp[0][i]
		}
	}
	return best
}

func main() {
	fmt.Println(minFallingPathSum([][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}))
}
