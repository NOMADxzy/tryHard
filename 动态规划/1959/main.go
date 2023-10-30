package main

import "fmt"

func minSpaceWastedKResizing(nums []int, k int) int {
	m := len(nums)
	if k+1 >= m {
		return 0
	}

	dpLeft, dpRight := make([]int, m), make([]int, m)
	dpLeft[0], dpRight[0] = nums[0], nums[m-1]
	sums := nums[0]

	for i := 1; i < m; i++ {
		sums += nums[i]
		i_ := m - 1 - i
		dpLeft[i] = dpLeft[i-1]
		if nums[i] > dpLeft[i] {
			dpLeft[i] = nums[i]
		}
		dpRight[i_] = dpRight[i_+1]
		if nums[i_] > dpRight[i_] {
			dpRight[i_] = nums[i_]
		}
	}

	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[0][j] = dpLeft[j] * (j + 1)
		}
	}

	for curK := 1; curK <= k; curK++ {
		for j := curK; j < m; j++ {
			best := dp[curK-1][j-1] + nums[j]
			rightMax := nums[j]
			for c := j - 1; c >= curK; c-- {
				if nums[c] > rightMax {
					rightMax = nums[c]
				}
				if val := rightMax*(j-c+1) + dp[curK-1][c-1]; val < best {
					best = val
				}
			}
			dp[curK][j] = best
		}
	}
	return dp[k][m-1] - sums
}

func main() {
	fmt.Println(minSpaceWastedKResizing([]int{10, 20, 30}, 0))
}
