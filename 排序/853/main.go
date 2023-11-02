package main

import (
	"fmt"
	"sort"
)

func carFleet(target int, position []int, speed []int) int {
	m := len(position)

	pair := make([][]int, m)
	for i := 0; i < m; i++ {
		pair[i] = []int{position[i], speed[i]}
	}

	sort.Slice(pair, func(i, j int) bool {
		return pair[i][0] < pair[j][0]
	})

	limitTime := float64(target-pair[m-1][0]) / float64(pair[m-1][1])
	ans := 1

	for i := m - 2; i >= 0; i-- {
		curTime := float64(target-pair[i][0]) / float64(pair[i][1])
		if curTime > limitTime {
			ans++
			limitTime = curTime
		}
	}
	return ans
}

func main() {
	fmt.Println(carFleet(17, []int{8, 12, 16, 11, 7}, []int{6, 9, 10, 9, 7}))
}
