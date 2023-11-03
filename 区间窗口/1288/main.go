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
		} else if lastOne >= 0 && intervals[i][0] == intervals[lastOne][0] {
			ans++
		}
		if intervals[i][1] > rightMax {
			rightMax = intervals[i][1]
			lastOne = i
		}
	}
	return len(intervals) - ans
}

func main() {
	fmt.Println(removeCoveredIntervals([][]int{{1, 3}, {2, 4}, {2, 3}, {2, 5}}))
}
