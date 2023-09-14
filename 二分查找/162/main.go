package main

import "fmt"

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		if l+1 == r {
			if nums[l] < nums[r] {
				return r
			} else {
				return l
			}
		} else {
			mid := (l + r) / 2
			if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
				return mid
			} else if nums[mid] < nums[mid-1] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return l
}

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
}
