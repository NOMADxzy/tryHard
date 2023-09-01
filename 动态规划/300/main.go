package main

import "fmt"

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	max := 1

	for i := 1; i < len(nums); i++ {
		longest := 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && longest < dp[j]+1 {
				longest = dp[j] + 1
			}
		}
		dp[i] = longest
		if longest > max {
			max = longest
		}
	}
	return max
}

func main() {
	fmt.Println(lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
}
