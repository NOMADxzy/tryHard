package main

import "fmt"

func sortColors(nums []int) {
	l, r := 0, len(nums)-1

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			l++
		} else if nums[i] == 2 {
			r--
		}
	}
	for i := 0; i < len(nums); i++ {
		if i < l {
			nums[i] = 0
		} else if i > r {
			nums[i] = 2
		} else {
			nums[i] = 1
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
