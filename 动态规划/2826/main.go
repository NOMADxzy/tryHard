package main

import "fmt"

func minimumOperations(nums []int) int {
	dp := []int{0, 0, 0}
	if nums[0] == 2 {
		dp[0]++
	} else if nums[0] == 3 {
		dp[0]++
		dp[1]++
	}

	for i := 1; i < len(nums); i++ {
		newDp := make([]int, 3)
		newDp[0] = dp[0]
		newDp[1] = dp[1]
		newDp[2] = dp[2]
		if nums[i] == 1 {
			newDp[0] = dp[0]
			newDp[1] = min(dp[0], dp[1]+1)
			newDp[2] = min(dp[0], dp[2]+1)

		} else if nums[i] == 2 {
			newDp[0] = dp[0] + 1
			newDp[1] = dp[1]
			newDp[2] = min(dp[1], dp[2]+1)
		} else {
			newDp[0] = dp[0] + 1
			newDp[1] = dp[1] + 1
			newDp[2] = dp[2]
		}
		dp = newDp
	}
	return min(dp[0], min(dp[1], dp[2]))
}

func main() {
	fmt.Println(minimumOperations([]int{2, 1, 3, 2, 1}))
}
