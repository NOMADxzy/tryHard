package main

import "fmt"

func getMoneyAmount(n int) int {
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}
	for l := 1; l < n; l++ {
		for i := 0; i < n-l; i++ {
			best := 1<<31 - 1
			for mid := i; mid <= i+l; mid++ {
				val := mid + 1
				left, right := 0, 0
				if mid > i {
					left = dp[i][mid-1]
				}
				if mid < i+l {
					right = dp[mid+1][i+l]
				}
				val += max(left, right)
				if val < best {
					best = val
				}
			}
			dp[i][i+l] = best
		}
	}
	return dp[0][n-1]
}

func main() {
	fmt.Println(getMoneyAmount(10))
}
