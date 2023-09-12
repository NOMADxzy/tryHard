package main

import (
	"fmt"
	"sort"
)

func maxOperations(nums []int, k int) int {
	cnt := 0
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	l := 0

	for i := len(nums) - 1; i >= 0; i-- {
		target := k - nums[i]

		for l < i && nums[l] < target {
			l++
		}
		if l >= i {
			break
		}

		if nums[l] == target {
			cnt += 1
			l++
		}

	}
	return cnt
}

func main() {
	fmt.Println(maxOperations([]int{3, 1, 3, 4, 3}, 6))
}
