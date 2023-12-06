package main

import "fmt"

func findMin(nums []int) int {
	m := len(nums)
	left, right := 0, m-1
	minVal := nums[0]
	for left < right {
		minVal = min(minVal, nums[left])
		mid := (left + right) / 2
		if nums[mid] > nums[left] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			if nums[mid] == nums[left] {
				left++
			}
			if nums[mid] == nums[right] {
				right--
			}
		}
	}
	return min(minVal, nums[right])
}

func main() {
	fmt.Println(findMin([]int{2, 2, 2, 0, 1}))
}
