package main

import (
	"fmt"
	"sort"
)

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	m := len(reward1)
	deltas := make([][2]int, m)
	for i := 0; i < m; i++ {
		deltas[i][0] = reward1[i] - reward2[i]
		deltas[i][1] = i
	}
	sort.Slice(deltas, func(i, j int) bool {
		return deltas[i][0] > deltas[j][0]
	})
	ans := 0
	for i := 0; i < k; i++ {
		ans += reward1[deltas[i][1]]
	}
	for i := k; i < m; i++ {
		ans += reward2[deltas[i][1]]
	}
	return ans
}

func main() {
	fmt.Println(miceAndCheese([]int{1, 1, 3, 4}, []int{4, 4, 1, 1}, 2))
}
