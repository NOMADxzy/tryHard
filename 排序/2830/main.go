package main

import (
	"fmt"
	"sort"
)

func maximizeTheProfit(n int, offers [][]int) int {
	sort.Slice(offers, func(i, j int) bool {
		return offers[i][1] < offers[j][1]
	})
	m := len(offers)
	dp := make([]int, m)
	for i := 0; i < m; i++ {
		if i > 0 {
			dp[i] = dp[i-1]
		}
		var j int
		if offers[0][1] >= offers[i][0] {
			j = -1
		} else {
			left, right := 0, i-1
			for left < right {
				mid := (left + right + 1) / 2
				if offers[mid][1] < offers[i][0] {
					left = mid
				} else {
					right = mid - 1
				}
			}
			j = left
		}
		val := offers[i][2]
		if j >= 0 {
			val += dp[j]
		}
		dp[i] = max(dp[i], val)
	}
	return dp[m-1]
}

func main() {
	fmt.Println(maximizeTheProfit(5, [][]int{{0, 0, 1}, {0, 2, 2}, {1, 3, 2}}))
}
