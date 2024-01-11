package main

import "fmt"

func minimumSize(nums []int, maxOperations int) int {
	left, right := 1, 1
	for i := 0; i < len(nums); i++ {
		right = max(right, nums[i])
	}
	getSteps := func(limit int) int {
		cnt := 0
		for i := 0; i < len(nums); i++ {
			cnt += nums[i] / limit
			if nums[i]%limit == 0 {
				cnt--
			}
		}
		return cnt
	}
	if getSteps(left) <= maxOperations {
		return left
	}
	for left < right {
		mid := (left + right) / 2
		if getSteps(mid) > maxOperations {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}

func main() {
	fmt.Println(minimumSize([]int{2, 4, 8, 2}, 4))
}
