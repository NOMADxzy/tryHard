package main

import "fmt"

func countPairs(n int, edges [][]int) int64 {
	parents := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
	}

	var find func(i int) int
	find = func(i int) int {
		if parents[i] != i {
			parents[i] = find(parents[i])
		}
		return parents[i]
	}
	union := func(i1, i2 int) {
		parents[find(i1)] = find(i2)
	}

	groups := map[int]int{}
	for _, edge := range edges {
		union(edge[0], edge[1])
	}
	for i := 0; i < n; i++ {
		groups[find(i)]++
	}
	ans := int64(0)
	for _, gsize := range groups {
		ans += int64(gsize * (n - gsize))
	}
	return ans / 2
}

func main() {
	fmt.Println(countPairs(7, [][]int{{0, 2}, {0, 5}, {2, 4}, {1, 6}, {5, 4}}))
}
