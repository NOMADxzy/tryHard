package main

import "fmt"

func putIn(ans []int, t int) {
	if t <= ans[2] {
		return
	} else if t < ans[1] {
		ans[2] = t
	} else if t == ans[1] {
		return
	} else if t < ans[0] {
		ans[2] = ans[1]
		ans[1] = t
	} else if t == ans[0] {
		return
	} else {
		ans[2] = ans[1]
		ans[1] = ans[0]
		ans[0] = t
	}
}

func getBiggestThree(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	dp := make([][][]int, m+2)
	for i := 0; i < m+2; i++ {
		dp[i] = make([][]int, n+2)
		for j := 0; j < n+2; j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			trueI, trueJ := i-1, j-1
			dp[i][j][0] = grid[trueI][trueJ] + dp[i-1][j-1][0]
			dp[i][j][1] = grid[trueI][trueJ] + dp[i-1][j+1][1]
		}
	}

	ans3 := make([]int, 3)

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			trueI, trueJ := i-1, j-1
			putIn(ans3, grid[trueI][trueJ])
			for d := 1; ; d++ {
				leftx, lefty, rightx, righty, bottomx, bottomy := i+d, j-d, i+d, j+d, i+2*d, j
				if leftx < 1 || leftx > m || lefty < 1 || lefty > n || rightx < 1 || rightx > m || righty < 1 || righty > n ||
					bottomx < 1 || bottomx > m || bottomy < 1 || bottomy > n {
					break
				}
				sum := dp[leftx][lefty][1] - dp[leftx-d][lefty+d][1] + dp[bottomx][bottomy][0] - dp[bottomx-d][bottomy-d][0] +
					dp[bottomx-1][bottomy+1][1] - dp[bottomx-d][bottomy+d][1] + dp[rightx][righty][0] - dp[rightx-d-1][righty-d-1][0]
				putIn(ans3, sum)
			}
		}
	}
	last := 2
	for last >= 0 && ans3[last] == 0 {
		last--
	}

	return ans3[:last+1]
}

func main() {
	grid := [][]int{{3, 4, 5, 1, 3}, {3, 3, 4, 2, 3}, {20, 30, 2000, 40, 10}, {1, 5, 5, 4, 1}, {4, 3, 2, 2, 5}}
	fmt.Println(getBiggestThree(grid))
}
