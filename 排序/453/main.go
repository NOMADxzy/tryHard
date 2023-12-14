package main

import (
	"fmt"
	"sort"
)

func minMoves(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	minVal := nums[0]
	ans := 0
	for i := 1; i < len(nums); i++ {
		if val := nums[i] - minVal; val > 0 {
			ans += val
		}
	}
	return ans
}

func main() {
	fmt.Println(minMoves([]int{1, 2, 3}))
}
