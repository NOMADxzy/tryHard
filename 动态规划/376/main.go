package main

import "fmt"

func max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func wiggleMaxLength(nums []int) int {
	m := len(nums)
	up := make([]int, m)
	down := make([]int, m)

	up[0] = 1
	down[0] = 1

	for i := 1; i < m; i++ {
		if nums[i] < nums[i-1] {
			down[i] = max(up[i-1]+1, down[i-1])
		} else {
			down[i] = down[i-1]
		}
		if nums[i] > nums[i-1] {
			up[i] = max(down[i-1]+1, up[i-1])
		} else {
			up[i] = up[i-1]
		}
	}
	return max(up[m-1], down[m-1])
}

func main() {
	fmt.Println(wiggleMaxLength([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}
