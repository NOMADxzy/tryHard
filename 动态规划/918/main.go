package main

import "fmt"

func maxSubarraySumCircular(nums []int) int {
	m := len(nums)
	dpMax := make([]int, m)
	dpMin := make([]int, m)
	sumAll := 0
	maxSum := nums[0]
	dpMax[0] = nums[0]
	dpMin[0] = nums[0]

	for i := 0; i < m; i++ {
		sumAll += nums[i]
	}

	for i := 1; i < m; i++ {
		if dpMax[i-1] > 0 {
			dpMax[i] = dpMax[i-1] + nums[i]
		} else {
			dpMax[i] = nums[i]
		}
		if dpMin[i-1] < 0 {
			dpMin[i] = dpMin[i-1] + nums[i]
		} else {
			dpMin[i] = nums[i]
		}
		if dpMax[i] > maxSum {
			maxSum = dpMax[i]
		}
		if i < m-1 && sumAll-dpMin[i] > maxSum {
			maxSum = sumAll - dpMin[i]
		}
	}

	return maxSum
}

func main() {
	fmt.Println(maxSubarraySumCircular([]int{-3, -2, -3}))
}
