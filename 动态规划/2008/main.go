package main

import (
	"fmt"
	"sort"
)

func maxTaxiEarnings(n int, rides [][]int) int64 {
	m := len(rides)
	sort.Slice(rides, func(i, j int) bool {
		return rides[i][1] < rides[j][1]
	})

	dp := make([]int64, m)
	dp[0] = int64(rides[0][2] + rides[0][1] - rides[0][0])
	for i := 1; i < m; i++ {
		dp[i] = dp[i-1]
		limit := rides[i][0]
		left, right := 0, i-1
		if rides[0][1] > limit {
			left = -1
		} else if rides[right][1] <= limit {
			left = right
		} else {
			for left < right {
				mid := (left + right) / 2
				if rides[mid][1] <= limit {
					left = mid + 1
				} else {
					right = mid
				}
			}
			left--
		}
		pre := int64(0)
		if left >= 0 {
			pre += dp[left]
		}
		if val := pre + int64(rides[i][2]+rides[i][1]-rides[i][0]); val > dp[i] {
			dp[i] = val
		}
	}
	return dp[m-1]
}

func main() {
	fmt.Println(maxTaxiEarnings(20, [][]int{{1, 6, 1}, {3, 10, 2}, {10, 12, 3}, {11, 12, 2}, {12, 15, 2}, {13, 18, 1}}))
}
