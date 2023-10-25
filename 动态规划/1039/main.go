package main

import "fmt"

func minScoreTriangulation(values []int) int {
	m := len(values)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
	}
	MAXINT := 10000000000

	for l := 3; l <= m; l++ {
		for i := 0; i <= m-l; i++ {
			left, right := i, i+l-1
			best := MAXINT
			for k := left + 1; k < right; k++ {
				if val := dp[left][k] + dp[k][right] + values[left]*values[right]*values[k]; val < best {
					best = val
				}
			}
			dp[left][right] = best
		}
	}
	return dp[0][m-1]
}

func main() {
	fmt.Println(minScoreTriangulation([]int{1, 3, 1, 4, 1, 5}))
}
