package main

import (
	"fmt"
	"sort"
)

func subSort(array []int) []int {
	l := len(array)
	if l == 0 {
		return []int{-1, -1}
	}
	idxs := make([]int, l)
	for i := 0; i < l; i++ {
		idxs[i] = i
	}
	sort.SliceStable(idxs, func(i, j int) bool {
		return array[idxs[i]] < array[idxs[j]]
	})
	m, n := 0, l-1
	for m < l && idxs[m] == m {
		m++
	}
	for n > m && idxs[n] == n {
		n--
	}
	if m >= n {
		m, n = -1, -1
	}
	return []int{m, n}
}

func main() {
	fmt.Println(subSort([]int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}))
}
