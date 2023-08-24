package main

import "fmt"

// 暴力方法
func maxSubArray1(nums []int) int {
	dp := make([][]int, len(nums))
	for i, _ := range dp {
		dp[i] = make([]int, len(nums))
	}
	max := -10000

	for l := 1; l <= len(nums); l++ {
		for i := 0; i <= len(nums)-l; i++ {
			if l == 1 {
				dp[i][i+l-1] = nums[i]
			} else {
				dp[i][i+l-1] = dp[i][i+l-2] + nums[i+l-1]
			}
			if dp[i][i+l-1] > max {
				max = dp[i][i+l-1]
			}
		}
	}
	return max
}

func maxSubArray(nums []int) int {
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func main() {
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
}
