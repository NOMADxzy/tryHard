package main

import "fmt"

func largestCombination(candidates []int) int {
	dp := make([]int, 32)
	for i := 0; i < len(candidates); i++ {
		newDp := make([]int, 32)
		cur := candidates[i]
		mask := 1
		copy(newDp, dp)
		for j := 0; mask <= cur; j++ {
			if cur&mask != 0 {
				newDp[j]++
			}
			mask *= 2
		}
		dp = newDp
	}
	ans := dp[0]
	for i := 0; i < len(dp); i++ {
		ans = max(ans, dp[i])
	}
	return ans
}

func main() {
	fmt.Println(largestCombination([]int{1, 3, 5, 9}))
}
