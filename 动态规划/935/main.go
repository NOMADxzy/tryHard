package main

import "fmt"

func knightDialer(n int) int {
	R, C := 4, 3
	MOD := 1000000007
	dp, dpNext := make([][]int, R), make([][]int, R)
	for i := 0; i < R; i++ {
		dp[i] = make([]int, C)
		dpNext[i] = make([]int, C)
		for j := 0; j < C; j++ {
			dp[i][j] = 1
		}
		if i == 3 {
			dp[i][0], dp[i][2] = 0, 0
		}
	}

	dirs := [][]int{{-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}}

	for epoch := 1; epoch < n; epoch++ {
		for r := 0; r < 3; r++ {
			for c := 0; c < 3; c++ {
				valids := 0
				for k := 0; k < 8; k++ {
					nextX, nextY := r+dirs[k][0], c+dirs[k][1]
					if (nextX >= 0 && nextX < 3 && nextY >= 0 && nextY < 3) || nextX == 3 && nextY == 1 {
						valids = (valids + dp[nextX][nextY]) % MOD
					}
				}
				dpNext[r][c] = valids
			}
		}
		dpNext[3][1] = (dp[1][0] + dp[1][2]) % MOD
		dp, dpNext = dpNext, dp
	}
	ans := dp[3][1]
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			ans = (dp[i][j] + ans) % MOD
		}
	}
	return ans
}

func main() {
	fmt.Println(knightDialer(3131))
}
