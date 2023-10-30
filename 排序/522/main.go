package main

import (
	"fmt"
	"sort"
)

func check(a, b string) bool {
	m, n := len(a), len(b)
	i, j := 0, 0
	if m > n {
		return false
	}
	for i < m && j < n {
		if a[i] == b[j] {
			i++
			j++
		} else {
			j++
		}
	}
	if i == m {
		return true
	}
	return false
}

func findLUSlength(strs []string) int {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	m := len(strs)
	//mark := make([]bool, m)
	if len(strs[m-1]) > len(strs[m-2]) {
		return len(strs[m-1])
	}
	for i := m - 1; i >= 0; i-- {
		invalid := false
		var j int
		for j = i; j > 0 && len(strs[j]) == len(strs[i]); j-- {
		}
		for ; j < m; j++ {
			if j != i {
				if check(strs[i], strs[j]) {
					invalid = true
				}
			}
		}
		if !invalid {
			return len(strs[i])
		}
	}
	return -1
}

func main() {
	fmt.Println(findLUSlength([]string{"aabbcc", "aabbcc", "c", "e", "aabbcd"}))
}
