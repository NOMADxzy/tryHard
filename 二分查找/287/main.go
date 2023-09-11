package main

import "fmt"

func findDuplicate(nums []int) int {
	min := 1
	max := len(nums) - 1

	for min < max {
		mid := (min + max) / 2
		cnt := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] > mid && nums[i] <= max {
				cnt++
			}
		}
		if cnt > max-mid {
			min = mid + 1
		} else {
			max = mid
		}
	}
	return min
}

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 1}))
}
