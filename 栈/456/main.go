package main

import "fmt"

func find132pattern(nums []int) bool {
	var stack []int

	k := -(1000000007)

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < k {
			return true
		}
		for len(stack) > 0 && stack[len(stack)-1] < nums[i] {
			if stack[len(stack)-1] > k {
				k = stack[len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false
}

func main() {
	fmt.Println(find132pattern([]int{3, 1, 4, 2}))
}
