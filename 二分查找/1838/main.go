package main

import (
	"fmt"
	"sort"
)

func maxFrequency(nums []int, k int) int {
	m := len(nums)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	sums := make([]int, m+1)
	for i := 1; i <= m; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}

	pre := 0
	best := 0
	for i := 0; i < m; i++ {
		cur := 0
		if i > 0 && nums[i] == nums[i-1] {
			cur = pre + 1
		} else {
			left, right := 0, i
			if nums[i]*(i)-sums[i] <= k {
				right = 0
			} else {
				for left < right {
					mid := (left + right) / 2
					if nums[i]*(i-mid)-sums[i]+sums[mid] <= k {
						right = mid
					} else {
						left = mid + 1
					}
				}
			}
			cur = i - right + 1
		}
		if cur > best {
			best = cur
		}
		pre = cur
	}
	return best
}

func main() {
	fmt.Println(maxFrequency([]int{1, 1, 2, 4}, 4))
}
