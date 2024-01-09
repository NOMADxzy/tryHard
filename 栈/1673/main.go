package main

import "fmt"

func mostCompetitive(nums []int, k int) []int {
	var stack []int
	for i := 0; i < len(nums); i++ {
		if len(stack) == 0 {
			stack = append(stack, nums[i])
		} else if nums[i] >= stack[len(stack)-1] {
			if len(stack) < k {
				stack = append(stack, nums[i])
			}
		} else {
			for len(stack) > 0 && stack[len(stack)-1] > nums[i] && len(nums)-i+len(stack)-1 >= k {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, nums[i])
		}
	}
	return stack
}

func main() {
	fmt.Println(mostCompetitive([]int{3, 5, 2, 6, 1}, 4))
}
