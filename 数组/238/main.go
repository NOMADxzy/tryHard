package main

import "fmt"

func productExceptSelf(nums []int) []int {
	m := len(nums)
	rightAccMulti := make([]int, m+1)
	rightAccMulti[m] = 1
	leftAccMulti := nums[0]

	for i := m - 1; i > 0; i-- {
		rightAccMulti[i] = rightAccMulti[i+1] * nums[i]
	}

	result := make([]int, m)
	result[0] = rightAccMulti[1]

	for i := 1; i < m; i++ {
		result[i] = leftAccMulti * rightAccMulti[i+1]
		leftAccMulti *= nums[i]
	}
	return result
}

func main() {
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}
