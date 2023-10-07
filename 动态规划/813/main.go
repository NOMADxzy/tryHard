package main

import "fmt"

func largestSumOfAverages(nums []int, k int) float64 {
	m := len(nums)
	sums := make([]float64, m)
	sums[0] = float64(nums[0])

	for i := 1; i < m; i++ {
		sums[i] = sums[i-1] + float64(nums[i])
	}

	dp := make([][]float64, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]float64, k+1)
	}
	for i := 0; i < m; i++ {
		dp[i][1] = sums[i] / float64(i+1)
	}

	for i := 1; i < m; i++ {
		for j := 2; j <= k && j <= i+1; j++ {
			max := float64(0)
			for ri := i - 1; ri >= j-2; ri-- {
				left := dp[ri][j-1]
				right := (sums[i] - sums[ri]) / float64(i-ri)
				if left+right > max {
					max = left + right
				}
			}
			dp[i][j] = max
		}
	}
	return dp[m-1][k]
}

func main() {
	fmt.Println(largestSumOfAverages([]int{9, 1, 2, 3, 9}, 3))
}
