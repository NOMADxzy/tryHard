package main

import (
	"fmt"
	"sort"
)

func findRightInterval(intervals [][]int) []int {
	m := len(intervals)
	idxs := make([]int, m)
	intervalsStart := make([][]int, m)
	for i := 0; i < m; i++ {
		intervalsStart[i] = make([]int, 3)
		intervalsStart[i] = append(intervals[i], i)
		intervals[i] = append(intervals[i], i)
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	sort.Slice(intervalsStart, func(i, j int) bool {
		return intervalsStart[i][0] < intervalsStart[j][0]
	})

	last := 0
	for i := 0; i < m; i++ {
		idx := intervals[i][2]

		if last == m {
			idxs[idx] = -1
		} else {
			for last < m && intervalsStart[last][0] < intervals[i][1] {
				last++
			}
			if last < m {
				idxs[idx] = intervalsStart[last][2]
			} else {
				idxs[idx] = -1
			}
		}
	}
	return idxs
}

func main() {
	fmt.Println(findRightInterval([][]int{{3, 4}, {2, 3}, {1, 2}}))
}
