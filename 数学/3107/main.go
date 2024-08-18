package main

import (
	"fmt"
	"sort"
)

//	10
//        5
// 1 2 3 4   7 8

// 2 5 5 6 8
// 17 22 40 74 75 86
// 1 2 3 4 5 6
func minOperationsToMakeMedianK(nums []int, k int) int64 {
	m := len(nums)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ans := int64(0)
	var i int
	for i = 0; i < m && nums[i] <= k; i++ {
	}
	if i == m {
		for j := m - 1; j >= m/2; j-- {
			ans += int64(k - nums[j])
		}
		return ans
	}
	if m%2 == 0 {
		if i <= m/2 {
			rightPos := m / 2
			for j := i; j <= rightPos; j++ {
				ans += int64(nums[j] - k)
			}
		} else {
			leftPos := m / 2
			for j := i - 1; j >= leftPos; j-- {
				ans += int64(k - nums[j])
			}
		}
	} else {
		if i <= m/2 {
			rightPos := m / 2
			for j := i; j <= rightPos; j++ {
				ans += int64(nums[j] - k)
			}
		} else {
			leftPos := m / 2
			for j := i - 1; j >= leftPos; j-- {
				ans += int64(k - nums[j])
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(minOperationsToMakeMedianK([]int{1, 2, 3, 4, 5, 6}, 4))
}
