package main

import (
	"fmt"
	"sort"
)

func wiggleSort(nums []int) {
	m := len(nums)
	arr := make([]int, m)
	copy(arr, nums)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	idx := 0
	for i := 1; i < m; i += 2 {
		nums[i] = arr[idx]
		idx++
	}

	for i := 0; i < m; i += 2 {
		nums[i] = arr[idx]
		idx++
	}
}

func main() {
	nums := []int{1, 5, 1, 1, 6, 4}
	wiggleSort(nums)
	fmt.Println(nums)
}
