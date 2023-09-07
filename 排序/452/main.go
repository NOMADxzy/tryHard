package main

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	m := len(points)
	arrows := 1
	interval := []int{points[0][0], points[0][1]}

	for i := 1; i < m; i++ {
		cur := points[i]
		if cur[0] > interval[1] {
			interval[0] = cur[0]
			interval[1] = cur[1]
			arrows++
		} else {
			interval[0] = cur[0]
			if interval[1] > cur[1] {
				interval[1] = cur[1]
			}
		}
	}

	return arrows
}

func main() {
	fmt.Println(findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
	fmt.Println(findMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}))
}
