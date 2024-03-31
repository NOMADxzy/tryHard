package main

import (
	"fmt"
	"sort"
)

func wiggleSort(nums []int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	m := len(nums)
	mid := 0
	if m%2 == 0 {
		mid = m / 2
	} else {
		mid = m/2 + 1
	}
	lefts, rights := nums[:mid], nums[mid:]
	ans := make([]int, m)
	idx := 0
	for i := 0; i < len(lefts) && i < len(rights); i++ {
		ans[idx] = lefts[i]
		ans[idx+1] = rights[i]
		idx += 2
	}
	if idx < m {
		ans[idx] = lefts[len(lefts)-1]
	}
	for i := 0; i < m; i++ {
		nums[i] = ans[i]
	}
}

func main() {
	nums := []int{5, 3, 1, 2, 3}
	wiggleSort(nums)
	fmt.Println(nums)
}
