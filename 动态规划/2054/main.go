package main

import (
	"fmt"
	"sort"
)

func maxTwoEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})
	m := len(events)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0], dp[0][1] = events[0][2], events[0][2]

	for i := 1; i < m; i++ {
		newVal := 0
		start, _, val := events[i][0], events[i][1], events[i][2]
		dp[i][0] = max(dp[i-1][0], val)
		dp[i][1] = max(val, dp[i-1][1])
		if start <= events[0][1] {
			newVal = val
		} else if start > events[i-1][1] {
			newVal = dp[i-1][0] + val
		} else {
			left, right := 0, i-1
			for left < right {
				mid := (left + right) / 2
				if start <= events[mid][1] {
					right = mid
				} else {
					left = mid + 1
				}
			}
			newVal = dp[left-1][0] + val
		}
		dp[i][1] = max(dp[i-1][1], newVal)
	}
	return dp[m-1][1]
}

func main() {
	fmt.Println(maxTwoEvents([][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}}))
}
