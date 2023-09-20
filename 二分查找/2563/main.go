package main

import (
	"fmt"
	"sort"
)

func countFairPairs(nums []int, lower int, upper int) int64 {
	m := len(nums)
	cnt := int64(0)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < m-1; i++ {
		l := i + 1
		r := m
		if nums[l]+nums[i] > upper || nums[r-1]+nums[i] < lower {
			continue
		}
		for l < r {
			mid := (l + r) / 2
			if nums[mid] > upper-nums[i] {
				r = mid
			} else {
				l = mid + 1
			}
		}
		r_border := r
		l = i + 1
		r = m - 1
		for l < r {
			mid := (l + r) / 2
			if nums[mid] < lower-nums[i] {
				l = mid + 1
			} else {
				r = mid
			}
		}
		l_border := l - 1
		cnt += int64(r_border - l_border - 1)

	}

	return cnt
}

func main() {
	fmt.Println(countFairPairs([]int{1, 7, 9, 2, 5}, 3, 6))
}
