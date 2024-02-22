package main

import (
	"fmt"
	"sort"
)

func partitionArray(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ans := 1
	pre := nums[0]
	for i := 0; i < len(nums); {
		var j int
		for j = i + 1; j < len(nums) && nums[j]-pre <= k; j++ {
		}
		if j == len(nums) {
			break
		} else {
			ans++
			pre = nums[j]
			i = j
		}
	}
	return ans
}

func main() {
	fmt.Println(partitionArray([]int{3, 6, 1, 2, 5}, 2))
}
