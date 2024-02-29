package main

import (
	"fmt"
	"sort"
)

func longestSquareStreak(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	nexts := map[int]int{}
	ans := -1
	for i := 0; i < len(nums); i++ {
		nexts[nums[i]*nums[i]] = nexts[nums[i]] + 1
		ans = max(ans, nexts[nums[i]]+1)
	}
	if ans < 2 {
		return -1
	}
	return ans
}

func main() {
	fmt.Println(longestSquareStreak([]int{4, 3, 6, 16, 8, 2}))
}
