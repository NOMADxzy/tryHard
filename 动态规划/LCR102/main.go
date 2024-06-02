package main

import "fmt"

func findTargetSumWays(nums []int, target int) (count int) {
	var backtrack func(int, int)
	backtrack = func(index, sum int) {
		if index == len(nums) {
			if sum == target {
				count++
			}
			return
		}
		backtrack(index+1, sum+nums[index])
		backtrack(index+1, sum-nums[index])
	}
	backtrack(0, 0)
	return
}

func findTargetSumWays1(nums []int, target int) int {
	dp := make([]int, 2001)
	dp[1000-target] = 1
	for i := 0; i < len(nums); i++ {
		newDp := make([]int, 2001)
		for j := 0; j < len(dp); j++ {
			if dp[j] > 0 {
				if j-1000+target+nums[i] >= target-1000 && j-1000+target+nums[i] <= target+1000 {
					newDp[j-1000+target+nums[i]+1000-target] += dp[j]
				}
				if j-1000+target-nums[i] >= target-1000 && j-1000+target-nums[i] <= target+1000 {
					newDp[j-1000+target-nums[i]+1000-target] += dp[j]
				}
			}
		}
		dp = newDp
	}
	return dp[target+1000]
}

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}
