package main

import (
	"fmt"
	"sort"
)

func minMoves2(nums []int) int {
	m := len(nums)
	ans := 1 << 31
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	sums := make([]int, m+1)
	for i := 1; i <= m; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}

	for i := 0; i < m; i++ {
		left := i*nums[i] - sums[i]
		right := sums[m] - sums[i+1] - (m-1-i)*nums[i]
		if left+right < ans {
			ans = left + right
		}
	}
	return ans
}

func main() {
	fmt.Println(minMoves2([]int{1, 10, 2, 9}))
}
