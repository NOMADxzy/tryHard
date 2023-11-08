package main

import (
	"fmt"
	"sort"
)

func reductionOperations(nums []int) int {
	m := len(nums)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ans := 0
	for i := m - 2; i >= 0; i-- {
		if nums[i] == nums[i+1] {
		} else {
			ans += m - 1 - i
		}
	}
	return ans
}

func main() {
	fmt.Println(reductionOperations([]int{5, 1, 3}))
}
