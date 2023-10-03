package main

import (
	"fmt"
	"sort"
)

func triangleNumber(nums []int) int {
	m := len(nums)
	if m < 3 {
		return 0
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ans := 0
	for i := 0; i < m-2; i++ {
		for j := i + 1; j < m-1; j++ {
			k := nums[i] + nums[j]
			l, r := j+1, m-1
			if nums[l] >= k {
				continue
			} else if nums[r] < k {
				ans += m - j - 1
			} else {
				for l < r {
					mid := (l + r) / 2
					if nums[mid] < k {
						l = mid + 1
					} else {
						r = mid
					}
				}
				ans += l - j - 1
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(triangleNumber([]int{4, 2, 3, 4}))
}
