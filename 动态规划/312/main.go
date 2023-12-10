package main

import "fmt"

func maxCoins(nums []int) int {
	newNums := make([]int, len(nums)+2)
	m := len(newNums)
	newNums[0] = 1
	newNums[m-1] = 1
	for i := 1; i < m-1; i++ {
		newNums[i] = nums[i-1]
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
		if i > 0 && i < m-1 {
			dp[i-1][i+1] = newNums[i-1] * newNums[i] * newNums[i+1]
		}
	}

	for l := 3; l < m; l++ {
		for i := 0; i < m-l; i++ {
			best := 0
			for mid := i + 1; mid < i+l; mid++ {
				if val := dp[i][mid] + dp[mid][i+l] + newNums[i]*newNums[mid]*newNums[i+l]; val > best {
					best = val
				}
			}
			dp[i][i+l] = best
		}
	}
	return dp[0][m-1]
}

func main() {
	fmt.Println(maxCoins([]int{3, 1, 5, 8}))
}
