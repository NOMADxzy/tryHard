package main

import "fmt"

func removeElement(nums []int, val int) int {
	count := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			count += 1
		} else {
			nums[i-count] = nums[i]
		}
	}
	return len(nums) - count
}

func main() {
	arr := []int{0, 1, 2, 2, 3, 0, 4, 2}
	l := removeElement(arr, 2)
	for i := 0; i < l; i++ {
		fmt.Print(arr[i], " ")
	}
}
