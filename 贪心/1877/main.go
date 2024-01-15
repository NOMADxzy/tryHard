package main

import (
	"fmt"
	"sort"
)

func minPairSum(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ans := 0
	for i := 0; i < len(nums)/2; i++ {
		ans = max(ans, nums[i]+nums[len(nums)-1-i])
	}
	return ans
}

func main() {
	fmt.Println(minPairSum([]int{3, 5, 2, 3}))
}
