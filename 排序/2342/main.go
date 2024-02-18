package main

import (
	"fmt"
	"sort"
)

func maximumSum(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	hist := map[int]int{}
	getVal := func(x int) int {
		sum := 0
		for ; x > 0; x /= 10 {
			sum += x % 10
		}
		return sum
	}
	ans := -1
	for i := 0; i < len(nums); i++ {
		v := getVal(nums[i])
		if hist[v] > 0 {
			ans = max(ans, nums[i]+hist[v])
		} else {
			hist[v] = nums[i]
		}

	}
	return ans
}

func main() {
	fmt.Println(maximumSum([]int{368, 369, 307, 304, 384, 138, 90, 279, 35, 396, 114, 328, 251, 364, 300, 191, 438, 467, 183}))
}
