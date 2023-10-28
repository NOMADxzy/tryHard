package main

import (
	"fmt"
	"sort"
)

func bestTeamScore(scores []int, ages []int) int {
	m := len(ages)
	pairs := make([][]int, m)
	for i := 0; i < m; i++ {
		pairs[i] = make([]int, 2)
		pairs[i][0] = scores[i]
		pairs[i][1] = ages[i]
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1]*1000000+pairs[i][0] < pairs[j][1]*1000000+pairs[j][0]
	})

	dp := make([]int, m)
	dp[0] = pairs[0][0]
	ans := pairs[0][0]

	for i := 1; i < m; i++ {
		s := pairs[i][0]
		best := 0
		j := i - 1
		//fmt.Println(len(pairs))
		for ; j >= 0 && pairs[i][1] == pairs[j][1]; j-- {
			if dp[j] > best {
				best = dp[j]
			}
		}
		for ; j >= 0; j-- {
			if pairs[j][0] <= pairs[i][0] && dp[j] > best {
				best = dp[j]
			}
		}
		dp[i] = s + best
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

func main() {
	fmt.Println(bestTeamScore([]int{1, 3, 7, 3, 2, 4, 10, 7, 5}, []int{4, 5, 2, 1, 1, 2, 4, 1, 4}))
}
