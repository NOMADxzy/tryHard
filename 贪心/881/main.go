package main

import (
	"fmt"
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	m := len(people)
	idx := make([]int, m)
	for i := 0; i < m; i++ {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return people[idx[i]] > people[idx[j]]
	})
	r := m - 1
	ans := 0
	for i := 0; i <= r; i++ {
		if people[idx[i]]+people[idx[r]] <= limit {
			r--
		}
		ans++
	}
	return ans
}

func main() {
	fmt.Println(numRescueBoats([]int{3, 2, 2, 1}, 3))
}
