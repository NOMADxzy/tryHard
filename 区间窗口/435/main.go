package main

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	end := intervals[0][1]
	cnt := 0

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < end {
			cnt++
			if intervals[i][1] <= end {
				end = intervals[i][1]
			}
		} else {
			end = intervals[i][1]
		}
	}
	return cnt
}

func main() {
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
}
