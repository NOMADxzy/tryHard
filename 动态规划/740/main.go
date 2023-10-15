package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func deleteAndEarn(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var nArr, nCnts []int
	nArr = append(nArr, nums[0])
	nCnts = append(nCnts, 1)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			nCnts[len(nCnts)-1]++
		} else {
			nArr = append(nArr, nums[i])
			nCnts = append(nCnts, 1)
		}
	}
	dp := make([][]int, len(nArr))
	for i := 0; i < len(nArr); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = nArr[0] * nCnts[0]
	for i := 1; i < len(nArr); i++ {
		if nArr[i] == nArr[i-1]+1 {
			dp[i][1] = dp[i-1][0] + nArr[i]*nCnts[i]
		} else {
			dp[i][1] = nArr[i]*nCnts[i] + max(dp[i-1][0], dp[i-1][1])
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
	}
	return max(dp[len(nArr)-1][0], dp[len(nArr)-1][1])
}

func main() {
	fmt.Println(deleteAndEarn([]int{2, 2, 3, 3, 3, 4}))
}
