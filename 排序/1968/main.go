package main

import (
	"fmt"
	"sort"
)

func rearrangeArray(nums []int) []int {
	m := len(nums)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	//n1,n2 := m/2, m - m/2
	ans := make([]int, m)
	idx := 0
	for i := 0; i < m; i += 2 {
		ans[i] = nums[idx]
		idx++
	}
	for i := 1; i < m; i += 2 {
		ans[i] = nums[idx]
		idx++
	}
	return ans
}

func main() {
	fmt.Println(rearrangeArray([]int{1, 2, 3, 4, 5, 6}))
}
