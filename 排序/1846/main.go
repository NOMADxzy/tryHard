package main

import (
	"fmt"
	"sort"
)

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	m := len(arr)
	pre := 1
	for i := 1; i < m; i++ {
		if arr[i]-pre > 1 {
			pre = pre + 1
		} else {
			pre = arr[i]
		}
	}
	return pre
}

func main() {
	fmt.Println(maximumElementAfterDecrementingAndRearranging([]int{2, 2, 1, 2, 1}))
}
