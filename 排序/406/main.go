package main

import (
	"fmt"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	m := len(people)
	sortedPeople := make([][]int, m)

	sort.Slice(people, func(i, j int) bool {
		return people[i][0] < people[j][0]
	})

	for i := 0; i < m; i++ {
		p := people[i]
		idx := 0
		rank := 0
		for rank < p[1] {
			if sortedPeople[idx] == nil || sortedPeople[idx][0] == p[0] {
				rank++
			}
			idx++
		}
		for sortedPeople[idx] != nil {
			idx++
		}
		sortedPeople[idx] = []int{p[0], p[1]}
	}
	return sortedPeople
}

func main() {
	people := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	fmt.Println(reconstructQueue(people))
}
