package main

import "fmt"

func cherryPickup(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][n-1] = grid[0][0] + grid[0][n-1]
	ds := []int{-1, 0, 1}
	ans := 0
	for i := 1; i < m; i++ {
		dp1 := make([][]int, n)
		for i_ := 0; i_ < n; i_++ {
			dp1[i_] = make([]int, n)
		}
		for j1 := 0; j1 <= i && j1 < n; j1++ {
			for j2 := n - 1; j2 >= n-i-1 && j2 >= 0; j2-- {
				cur := grid[i][j1] + grid[i][j2]
				if j1 == j2 {
					cur /= 2
				}
				bestPre := 0
				for k1 := 0; k1 < 3; k1++ {
					prej1 := j1 + ds[k1]
					if prej1 < 0 || prej1 >= n || prej1 > i-1 {
						continue
					}
					for k2 := 0; k2 < 3; k2++ {
						prej2 := j2 + ds[k2]
						if prej2 < 0 || prej2 >= n || prej2 <= n-i-1 {
							continue
						}
						if dp[prej1][prej2] > bestPre {
							bestPre = dp[prej1][prej2]
						}
					}
				}
				dp1[j1][j2] = cur + bestPre
				ans = max(ans, cur+bestPre)
			}
		}
		dp = dp1
	}
	return ans
}

func main() {
	fmt.Println(cherryPickup([][]int{
		{3, 1, 1},
		{2, 5, 1},
		{1, 5, 5},
		{2, 1, 1},
	}))
}
