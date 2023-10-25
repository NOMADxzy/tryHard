package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func mctFromLeafValues(arr []int) int {
	m := len(arr)
	dp := make([][][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = make([]int, 2)
		}
		dp[i][i][1] = arr[i]
		if i < m-1 {
			dp[i][i+1][0] = arr[i] * arr[i+1]
			dp[i][i+1][1] = max(arr[i], arr[i+1])
		}
	}

	for l := 3; l <= m; l++ {
		for i := 0; i <= m-l; i++ {
			left, right := i, i+l-1
			best := dp[left+1][right][0] + arr[left]*dp[left+1][right][1]
			dp[left][right][1] = max(arr[left], dp[left+1][right][1])
			for k := left + 1; k < right; k++ {
				val := dp[left][k][0] + dp[k+1][right][0]
				val += dp[left][k][1] * dp[k+1][right][1]
				if val < best {
					best = val
				}
			}
			dp[left][right][0] = best
		}
	}
	return dp[0][m-1][0]
}

func main() {
	fmt.Println(mctFromLeafValues([]int{7, 12, 8, 10}))
}
