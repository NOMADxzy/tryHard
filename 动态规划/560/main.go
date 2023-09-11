package main

import "fmt"

func subarraySum(nums []int, k int) int {
	var dp []int

	dp = append(dp, nums[0])
	times := 0
	if dp[0] == k {
		times++
	}

	for i := 1; i < len(nums); i++ {
		dp = append(dp, 0)
		for j := 0; j < len(dp); j++ {
			dp[j] += nums[i]
			if dp[j] == k {
				times += 1
			}
		}

	}
	return times
}

func main() {
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
}
