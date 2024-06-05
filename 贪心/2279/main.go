package main

import "sort"

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	n := len(capacity)
	needs := make([]int, n)
	for i := 0; i < n; i++ {
		needs[i] = capacity[i] - rocks[i]
	}
	sort.Slice(needs, func(i, j int) bool {
		return needs[i] < needs[j]
	})
	ans := 0
	for i := 0; i < n; i++ {
		if needs[i] <= additionalRocks {
			additionalRocks -= needs[i]
			ans += 1
		} else {
			break
		}
	}
	return ans
}

func main() {
	maximumBags([]int{2, 3, 4, 5}, []int{1, 2, 4, 4}, 2)
}
