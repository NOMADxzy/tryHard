package main

import "fmt"

func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	m := len(source)
	parents := make([]int, m)
	for i := 0; i < m; i++ {
		parents[i] = i
	}

	var find func(parents []int, i int) int
	find = func(parents []int, i int) int {
		if parents[i] != i {
			parents[i] = find(parents, parents[i])
		}
		return parents[i]
	}

	union := func(parents []int, i1, i2 int) {
		parents[find(parents, i1)] = find(parents, i2)
	}

	groups := map[int]map[int]int{}
	target_group := map[int][]int{}

	for _, swap := range allowedSwaps {
		union(parents, swap[0], swap[1])
	}

	for i := 0; i < m; i++ {
		id := find(parents, i)
		target_group[id] = append(target_group[id], target[i])
		if groups[id] == nil {
			groups[id] = map[int]int{}
		}
		groups[id][source[i]]++
	}

	ans := 0
	for id, g := range target_group {
		sourceGroupMap := groups[id]
		for _, val := range g {
			if sourceGroupMap[val] > 0 {
				sourceGroupMap[val]--
			} else {
				ans++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(minimumHammingDistance([]int{5, 1, 2, 4, 3}, []int{1, 5, 4, 2, 3}, [][]int{{0, 4}, {4, 2}, {1, 3}, {1, 4}}))
}
