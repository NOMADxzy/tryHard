package main

import "fmt"

func orderOfLargestPlusSign(n int, mines [][]int) int {
	board := make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, n)
		for j := 0; j < n; j++ {
			board[i][j] = 1
		}
	}
	for _, mine := range mines {
		board[mine[0]][mine[1]] = 0
	}

	dp := make([][][]int, 4)
	for i := 0; i < 4; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, n)
		}
	}
	for i := 0; i < n; i++ {
		if board[0][i] == 1 {
			dp[0][0][i] = 1
		}
		if board[n-1][i] == 1 {
			dp[2][n-1][i] = 1
		}
		if board[i][0] == 1 {
			dp[3][i][0] = 1
		}
		if board[i][n-1] == 1 {
			dp[1][i][n-1] = 1
		}
	}
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 1 {
				dp[0][i][j] = dp[0][i-1][j] + 1
			}
			if board[n-1-i][j] == 1 {
				dp[2][n-1-i][j] = dp[2][n-i][j] + 1
			}
		}
	}
	for j := 1; j < n; j++ {
		for i := 0; i < n; i++ {
			if board[i][j] == 1 {
				dp[3][i][j] = dp[3][i][j-1] + 1
			}
			if board[i][n-1-j] == 1 {
				dp[1][i][n-1-j] = dp[1][i][n-j] + 1
			}
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			cur_ans := dp[0][i][j]
			if dp[1][i][j] < cur_ans {
				cur_ans = dp[1][i][j]
			}
			if dp[2][i][j] < cur_ans {
				cur_ans = dp[2][i][j]
			}
			if dp[3][i][j] < cur_ans {
				cur_ans = dp[3][i][j]
			}
			if cur_ans > ans {
				ans = cur_ans
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(orderOfLargestPlusSign(0, [][]int{}))
}
