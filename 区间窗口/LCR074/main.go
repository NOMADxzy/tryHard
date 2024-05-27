package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var ans [][]int
	for i := 0; i < len(intervals); i++ {
		start, end := intervals[i][0], intervals[i][1]
		j := i + 1
		for j < len(intervals) && intervals[j][0] <= end {
			end = max(end, intervals[j][1])
			j++
		}
		ans = append(ans, []int{start, end})
		i = j - 1
	}
	return ans
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}
