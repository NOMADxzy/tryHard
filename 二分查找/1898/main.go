package main

import "fmt"

func maximumRemovals(s string, p string, removable []int) int {
	idxs := map[int]int{}
	m := len(removable)
	for i := 0; i < m; i++ {
		idxs[removable[i]] = i + 1
	}
	check := func(k int) bool {
		i, j := 0, 0
		for i < len(s) && j < len(p) {
			for i < len(s) && idxs[i] <= k && idxs[i] > 0 {
				i++
			}
			if i == len(s) {
				break
			}
			if s[i] == p[j] {
				i++
				j++
			} else {
				i++
			}
		}
		return j == len(p)
	}
	left, right := 0, m
	if check(right) {
		return right
	} else {
		for left < right {
			mid := (left + right) / 2
			if check(mid) {
				left = mid + 1
			} else {
				right = mid
			}
		}
		return right - 1
	}
}

func main() {
	fmt.Println(maximumRemovals("abcacb", "ab", []int{3, 1, 0}))
}
