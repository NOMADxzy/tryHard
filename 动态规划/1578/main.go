package main

import "fmt"

func minCost(colors string, neededTime []int) int {
	m := len(colors)
	dp := make([]int, m)
	sums := make([]int, m)
	sums[0] = neededTime[0]
	for i := 1; i < m; i++ {
		sums[i] = sums[i-1] + neededTime[i]
	}

	preDifferentIdx := -1
	preSameSum := neededTime[0]
	for i := 1; i < m; i++ {
		if colors[i] != colors[i-1] {
			preDifferentIdx = i - 1
			preSameSum = neededTime[i]
			dp[i] = dp[i-1]
		} else {
			dp[i] = dp[i-1] + neededTime[i]

			dpPre := 0
			if preDifferentIdx > 0 {
				dpPre = dp[preDifferentIdx]
			}
			if dpPre+preSameSum < dp[i] {
				dp[i] = dpPre + preSameSum
			}
			preSameSum += neededTime[i]
		}
	}

	return dp[m-1]
}

func main() {
	fmt.Println(minCost("abaac", []int{1, 2, 3, 4, 5}))
}
