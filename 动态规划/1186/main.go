package main

import "fmt"

func maximumSum(arr []int) int {
	m := len(arr)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = arr[0]
	dp[0][1] = 0

	ans := arr[0]
	for i := 1; i < m; i++ {
		dp[i][0] = arr[i]
		if dp[i-1][0] > 0 {
			dp[i][0] += dp[i-1][0]
		}
		dp[i][1] = dp[i-1][0]
		if dp[i-1][1]+arr[i] > dp[i][1] {
			dp[i][1] = dp[i-1][1] + arr[i]
		}
		if arr[i] > dp[i][1] {
			dp[i][1] = arr[i]
		}

		if dp[i][0] > ans {
			ans = dp[i][0]
		}
		if dp[i][1] > ans {
			ans = dp[i][1]
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumSum([]int{1, -2, -2, 3}))
}
