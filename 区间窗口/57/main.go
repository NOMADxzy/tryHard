package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	var ans [][]int

	start, end := newInterval[0], newInterval[1]
	flag := false
	for i := 0; i < len(intervals); i++ {
		itv := intervals[i]
		if itv[1] < start || itv[0] > end {
			if !flag && itv[0] > end {
				flag = true
				ans = append(ans, []int{start, end})
			}
			ans = append(ans, itv)
			continue
		} else {
			if itv[0] < start {
				start = itv[0]
			}
			if itv[1] > end {
				end = itv[1]
			}
		}
	}
	if !flag {
		ans = append(ans, []int{start, end})
	}
	return ans
}

func main() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{0, 11}))
}
