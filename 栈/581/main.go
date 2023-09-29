package main

import "fmt"

func findUnsortedSubarray(nums []int) int {
	var stack []int
	right := 0
	findNextLeft := false
	left := right
	min := nums[len(nums)-1]
	for i := len(nums) - 1; i >= 0; i-- {
		if findNextLeft && nums[i] <= min {
			min = nums[i]
			left = i
			findNextLeft = false
		}
		if len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			if nums[stack[len(stack)-1]] < min {
				min = nums[stack[len(stack)-1]]
			}
			findNextLeft = true
			for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
				idx := stack[len(stack)-1]
				if idx > right {
					right = idx
				}
				stack = stack[:len(stack)-1]
			}
		}

		stack = append(stack, i)
	}
	if findNextLeft {
		left = -1
	}
	return right - left
}

func main() {
	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
}
