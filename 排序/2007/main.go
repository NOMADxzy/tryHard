package main

import (
	"fmt"
	"sort"
)

func findOriginalArray(changed []int) []int {
	m := len(changed)

	sort.Slice(changed, func(i, j int) bool {
		return changed[i] < changed[j]
	})

	cntMap := map[int]int{}
	for i := 0; i < m; i++ {
		cntMap[changed[i]]++
	}

	if m%2 != 0 {
		return []int{}
	}
	original := make([]int, m/2)
	idx := 0

	for i := 0; i < m; i++ {
		if cntMap[changed[i]] > 0 {
			if !(cntMap[2*changed[i]] > 0) {
				return []int{}
			} else {
				cntMap[changed[i]]--
				cntMap[2*changed[i]]--
				original[idx] = changed[i]
				idx++
			}
		}
	}
	return original
}

func main() {
	fmt.Println(findOriginalArray([]int{1, 3, 4, 2, 6, 8}))
}
