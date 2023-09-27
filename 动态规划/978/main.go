package main

import "fmt"

func maxTurbulenceSize(arr []int) int {
	maxLen := 1
	dp := make([][]int, len(arr))
	for i := 0; i < len(dp); i++ {
		dp[i] = []int{1, 1}
	}
	for i := 1; i < len(arr); i++ {
		if i%2 != 0 {
			if arr[i] > arr[i-1] {
				dp[i][0] = dp[i-1][0] + 1
			} else if arr[i] < arr[i-1] {
				dp[i][1] = dp[i-1][1] + 1
			}
		} else if i%2 == 0 {
			if arr[i] < arr[i-1] {
				dp[i][0] = dp[i-1][0] + 1
			} else if arr[i] > arr[i-1] {
				dp[i][1] = dp[i-1][1] + 1
			}
		}
		if dp[i][0] > maxLen {
			maxLen = dp[i][0]
		}
		if dp[i][1] > maxLen {
			maxLen = dp[i][1]
		}
	}
	return maxLen
}

func main() {
	fmt.Println(maxTurbulenceSize([]int{0, 1, 1, 0, 1, 0, 1, 1, 0, 0}))
}
