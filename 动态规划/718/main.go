package main

import "fmt"

func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	maxLen := 0
	for i := 0; i < m; i++ {
		if nums2[0] == nums1[i] {
			dp[i][0] = 1
			maxLen = 1
		}
	}
	for i := 0; i < n; i++ {
		if nums1[0] == nums2[i] {
			dp[0][i] = 1
			maxLen = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if nums1[i] == nums2[j] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = 0
			}
			if dp[i][j] > maxLen {
				maxLen = dp[i][j]
			}
		}
	}
	return maxLen
}

func main() {
	fmt.Println(findLength([]int{1, 2, 3, 2, 1, 4}, []int{3, 2, 1, 4, 7}))
}
