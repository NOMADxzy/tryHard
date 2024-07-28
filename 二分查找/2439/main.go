package main

import "fmt"

func minimizeArrayValue(nums []int) int {
	m := len(nums)
	maxVal := nums[0]
	for i := 0; i < m; i++ {
		maxVal = max(maxVal, nums[i])
	}
	check := func(limit int) bool {
		cnt := 0
		for i := 0; i < m; i++ {
			if nums[i] < limit {
				cnt += limit - nums[i]
			} else {
				cnt -= nums[i] - limit
				if cnt < 0 {
					return false
				}
			}
		}
		return true
	}
	left, right := 0, maxVal
	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func main() {
	fmt.Println(minimizeArrayValue([]int{3, 7, 1, 6}))
}
