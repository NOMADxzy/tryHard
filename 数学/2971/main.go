package main

import (
	"fmt"
	"sort"
)

func largestPerimeter(nums []int) int64 {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	acc := int64(nums[0] + nums[1])
	ans := int64(0)
	for i := 2; i < len(nums); i++ {
		if acc > int64(nums[i]) {
			ans = acc + int64(nums[i])
		}
		acc += int64(nums[i])
	}
	if ans == 0 {
		return -1
	}
	return ans
}

func main() {
	fmt.Println(largestPerimeter([]int{1, 12, 1, 2, 5, 50, 3}))
}
