package main

import (
	"fmt"
	"sort"
)

func maxNumOfMarkedIndices(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var i, j int
	ans := 0
	i, j = len(nums)/2-1, len(nums)-1
	for j >= len(nums)/2 && i >= 0 {
		for i >= 0 && nums[j] < 2*nums[i] {
			i--
		}
		if i < 0 {
			break
		}
		i--
		j--
		ans += 2
	}
	return ans
}

func main() {
	fmt.Println(maxNumOfMarkedIndices([]int{9, 2, 5, 4}))
}
