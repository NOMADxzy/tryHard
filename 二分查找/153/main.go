package main

import "fmt"

func find(nums []int, left int, right int) int {
	if left == right {
		return nums[left]
	}
	mid := (left + right) / 2
	if nums[mid] > nums[right] {
		return find(nums, mid+1, right)
	} else {
		return find(nums, left, mid)
	}
}

func findMin(nums []int) int {
	return find(nums, 0, len(nums)-1)
}

func main() {
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1}))
}
