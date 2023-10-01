package main

import (
	"fmt"
	"sort"
)

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})

	dp := make([]int, len(pairs))
	dp[0] = 1
	pre := 0
	for i := 1; i < len(pairs); i++ {
		cur := pairs[i]

		if cur[0] > pairs[pre][1] {
			dp[i] = dp[pre] + 1
			pre = i
		} else if cur[1] < pairs[pre][1] {
			dp[i] = dp[pre]
			pre = i
		}
	}
	return dp[pre]
}

func main() {
	fmt.Println(findLongestChain([][]int{{1, 2}, {2, 3}, {3, 4}}))
}
