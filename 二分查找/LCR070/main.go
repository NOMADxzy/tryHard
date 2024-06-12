package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	check := func(i int) int {
		if i == 0 {
			if nums[i] != nums[i+1] {
				return 0
			}
			return 1
		} else if i == len(nums)-1 {
			if nums[i] != nums[i-1] {
				return 0
			}
			return -1
		} else {
			if nums[i] != nums[i-1] && nums[i] != nums[i+1] {
				return 0
			} else if nums[i] == nums[i-1] {
				if (i-1)%2 == 0 {
					return 1
				} else {
					return -1
				}
			} else {
				if (len(nums)-1-(i+1))%2 == 0 {
					return -1
				} else {
					return 1
				}
			}
		}
	}
	left, right := 0, len(nums)-1
	if nums[0] != nums[1] {
		return nums[0]
	}
	if nums[len(nums)-1] != nums[len(nums)-2] {
		return nums[len(nums)-1]
	}
	for left < right {
		mid := (left + right) / 2
		v := check(mid)
		if v == 0 {
			return nums[mid]
		} else if v == 1 {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return -1
}

func main() {
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
}
