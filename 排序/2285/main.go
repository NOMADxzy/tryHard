package main

import (
	"fmt"
	"sort"
)

func maximumImportance(n int, roads [][]int) int64 {
	degrees := make([]int, n)
	for _, road := range roads {
		degrees[road[0]]++
		degrees[road[1]]++
	}
	sort.Slice(degrees, func(i, j int) bool {
		return degrees[i] > degrees[j]
	})
	ans := int64(0)
	for i := 0; i < n; i++ {
		ans += int64(degrees[i]) * int64(n-i)
	}
	return ans
}

func main() {
	roads := [][]int{{0, 3}, {2, 4}, {1, 3}}
	fmt.Println(maximumImportance(5, roads))
}
