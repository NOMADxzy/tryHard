package main

import "fmt"

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	if maxMove == 0 || (startRow >= maxMove && m-startRow > maxMove && startColumn >= maxMove && n-startColumn > maxMove) {
		return 0
	}
	MOD := 1000000007
	dp := make([][][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, maxMove)
			if i == 0 {
				dp[i][j][0] += 1
			}
			if i == m-1 {
				dp[i][j][0] += 1
			}
			if j == 0 {
				dp[i][j][0] += 1
			}
			if j == n-1 {
				dp[i][j][0] += 1
			}
		}
	}

	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for move := 1; move < maxMove; move++ {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				for d := 0; d < 4; d++ {
					nX, nY := dirs[d][0]+i, dirs[d][1]+j
					if nX >= 0 && nX < m && nY >= 0 && nY < n {
						dp[i][j][move] = (dp[i][j][move] + dp[nX][nY][move-1]%MOD) % MOD
					}
				}
			}
		}
	}
	ans := 0
	for i := 0; i < maxMove; i++ {
		ans = (ans + dp[startRow][startColumn][i]%MOD) % MOD
	}
	return ans
}

func main() {
	fmt.Println(findPaths(2, 2, 2, 0, 0))
}
