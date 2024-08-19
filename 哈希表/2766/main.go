package main

import (
	"fmt"
	"sort"
)

func relocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for i := 0; i < len(moveFrom); i++ {
		if moveFrom[i] == moveTo[i] { // TODO 易错，避免重复计数
			continue
		}
		m[moveTo[i]] += m[moveFrom[i]]
		m[moveFrom[i]] = 0
	}
	var ans []int
	for k, v := range m {
		if v > 0 {
			ans = append(ans, k)
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})
	return ans
}

func main() {
	fmt.Println(relocateMarbles([]int{1, 1, 3, 3}, []int{1, 3}, []int{2, 2}))
}
