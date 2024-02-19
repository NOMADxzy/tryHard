package main

import (
	"fmt"
	"sort"
)

func findScore(nums []int) int64 {
	score := int64(0)
	m := len(nums)
	mark := make([]bool, m)
	idxs := make([]int, m)
	for i := 0; i < m; i++ {
		idxs[i] = i
	}
	sort.Slice(idxs, func(i, j int) bool {
		return nums[idxs[i]] < nums[idxs[j]] || nums[idxs[i]] == nums[idxs[j]] && idxs[i] < idxs[j] //TODO 易错，
		// 快速排序是不稳定的，某些情况下需要加个判断使其稳定
	})
	for i := 0; i < m; i++ {
		if !mark[idxs[i]] {
			score += int64(nums[idxs[i]])
			if idxs[i]-1 >= 0 {
				mark[idxs[i]-1] = true
			}
			if idxs[i]+1 < m {
				mark[idxs[i]+1] = true
			}
		}
	}
	return score
}

func main() {
	fmt.Println(findScore([]int{2, 1, 3, 4, 5, 2}))
}
