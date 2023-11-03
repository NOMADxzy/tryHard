package main

import (
	"fmt"
	"sort"
)

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := 0
	rightMax := 0
	lastOne := -1
	for i := 0; i < len(intervals); i++ {
		if i > 0 && intervals[i][1] <= rightMax {
			ans++
			lastOne = i
		} else if i > 0 && intervals[i][0] == intervals[i-1][0] && i-1 != lastOne {
			ans++
		}
		if intervals[i][1] > rightMax {
			rightMax = intervals[i][1]
		}
	}
	return len(intervals) - ans
}

func main() {
	fmt.Println(removeCoveredIntervals([][]int{{1, 4}, {3, 6}, {2, 8}}))
}
