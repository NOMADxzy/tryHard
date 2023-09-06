package main

import "fmt"

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	minLeft := nums[0]
	maxRight := make([]int, len(nums))
	maxRight[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i > 1; i-- {
		if nums[i] > maxRight[i+1] {
			maxRight[i] = nums[i]
		} else {
			maxRight[i] = maxRight[i+1]
		}
	}

	for i := 1; i < len(nums)-1; i++ {
		if minLeft < nums[i] && nums[i] < maxRight[i+1] {
			return true
		} else {
			if nums[i] < minLeft {
				minLeft = nums[i]
			}
		}
	}
	return false
}

func main() {
	fmt.Println(increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
}
