package main

import (
	"fmt"
	"sort"
)

func maxDistance(position []int, m int) int {
	sort.Slice(position, func(i, j int) bool {
		return position[i] < position[j]
	})
	check := func(d int) bool {
		pre := position[0]
		cnt := 1
		for i := 1; i < len(position) && cnt < m; i++ {
			if position[i]-pre >= d {
				cnt++
				pre = position[i]
			}
		}
		return cnt == m
	}
	left, right := 1, position[len(position)-1]-position[0]/(m-1)
	if check(right) {
		return right
	}
	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right - 1
}

func main() {
	fmt.Println(maxDistance([]int{1, 2, 3, 4, 7}, 4))
}
