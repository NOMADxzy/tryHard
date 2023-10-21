package main

import "fmt"

func predictTheWinner(nums []int) bool {
	m := len(nums)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
		dp[i][i] = nums[i]
	}
	sums := make([]int, m)
	sums[0] = nums[0]
	for i := 1; i < m; i++ {
		sums[i] = sums[i-1] + nums[i]
	}

	for l := 2; l <= m; l++ {
		for start := 0; start <= m-l; start++ {
			end := l + start - 1
			score := nums[start] + (sums[end] - sums[start]) - dp[start+1][end]
			sum_start_1 := 0
			if start-1 >= 0 {
				sum_start_1 = sums[start-1]
			}
			if nums[end]+(sums[end-1]-sum_start_1)-dp[start][end-1] > score {
				score = nums[end] + (sums[end-1] - sum_start_1) - dp[start][end-1]
			}
			dp[start][end] = score
		}
	}
	return dp[0][m-1] >= sums[m-1]-dp[0][m-1]
}

func main() {
	fmt.Println(predictTheWinner([]int{1}))
}
