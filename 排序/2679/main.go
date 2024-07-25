package main

import (
	"fmt"
	"sort"
)

func matrixSum(nums [][]int) int {
	m := len(nums)
	n := len(nums[0])
	for i := 0; i < m; i++ {
		sort.Slice(nums[i], func(k, j int) bool {
			return nums[i][k] > nums[i][j]
		})
	}
	score := 0
	for j := 0; j < n; j++ {
		val := nums[0][j]
		for i := 1; i < m; i++ {
			if nums[i][j] > val {
				val = nums[i][j]
			}
		}
		score += val
	}
	return score
}

func main() {
	fmt.Println(matrixSum([][]int{{7, 2, 1}, {6, 4, 2}, {6, 5, 3}, {3, 2, 1}}))
}
