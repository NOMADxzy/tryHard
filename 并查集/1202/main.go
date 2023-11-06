package main

import (
	"fmt"
	"sort"
)

func smallestStringWithSwaps(s string, pairs [][]int) string {
	m := len(s)
	parents := make([]int, m)
	for i := 0; i < m; i++ {
		parents[i] = i
	}

	var find func(parents []int, index int) int
	find = func(parents []int, index int) int {
		if parents[index] != index {
			parents[index] = find(parents, parents[index])
		}
		return parents[index]
	}

	union := func(parents []int, i1, i2 int) {
		parents[find(parents, i1)] = find(parents, i2)
	}

	for _, pair := range pairs {
		union(parents, pair[0], pair[1])
	}
	trees := map[int][]int{}
	for i := 0; i < m; i++ {
		root := find(parents, i)
		trees[root] = append(trees[root], i)
	}
	ans := make([]byte, m)
	for _, idxs := range trees {
		sortedIdxs := make([]int, len(idxs))
		copy(sortedIdxs, idxs)

		sort.Slice(sortedIdxs, func(i, j int) bool {
			return s[sortedIdxs[i]] < s[sortedIdxs[j]]
		})

		for i := 0; i < len(idxs); i++ {
			ans[idxs[i]] = s[sortedIdxs[i]]
		}
	}
	return string(ans)
}

func main() {
	fmt.Println(smallestStringWithSwaps("dcab", [][]int{{0, 3}, {1, 2}}))
}
