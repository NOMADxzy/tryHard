package main

import "fmt"

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
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
		root1, root2 := find(i1), find(i2)
		parents[root1] = root2
	}
	for _, edge := range edges {
		v1, v2 := edge[0]-1, edge[1]-1
		root1, root2 := find(v1), find(v2)
		if root1 == root2 {
			return edge
		}
		union(v1, v2)
	}
	return nil
}

func main() {
	fmt.Println(findRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}}))
}
