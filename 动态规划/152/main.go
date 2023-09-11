package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func maxProduct(nums []int) int {
	dpPostive := make([]int, len(nums))
	dpNegative := make([]int, len(nums))

	dpPostive[0] = nums[0]
	dpNegative[0] = nums[0]
	m := nums[0]

	for i := 1; i < len(nums); i++ {
		dpPostive[i] = max(max(nums[i], dpPostive[i-1]*nums[i]), dpNegative[i-1]*nums[i])
		dpNegative[i] = min(min(nums[i], dpPostive[i-1]*nums[i]), dpNegative[i-1]*nums[i])
		if dpPostive[i] > m {
			m = dpPostive[i]
		}
	}
	return m
}

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4, -4}))
}
