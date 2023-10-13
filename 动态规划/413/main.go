package main

import "fmt"

func numberOfArithmeticSlices(nums []int) int {
	m := len(nums)
	dp := make([]int, m)
	ans := 0

	for i := 2; i < m; i++ {
		if nums[i-1]-nums[i-2] == nums[i]-nums[i-1] {
			dp[i] = dp[i-1] + 1
			ans += dp[i]
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3, 4}))
}
