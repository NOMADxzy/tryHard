package main

import "fmt"

func findCircleNum(isConnected [][]int) int {
	m := len(isConnected)
	parents := make([]int, m)
	for i := 0; i < m; i++ {
		parents[i] = i
	}

	var find func([]int, int) int
	find = func(parents []int, index int) int {
		if parents[index] != index {
			return find(parents, parents[index])
		} else {
			return index
		}
	}
	union := func(parents []int, index1, index2 int) {
		parents[find(parents, index1)] = find(parents, index2)
	}

	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			if isConnected[i][j] == 1 {
				union(parents, i, j)
			}
		}
	}

	provinceMap := map[int]bool{}
	ans := 0
	for idx := 0; idx < len(parents); idx++ {
		root := find(parents, idx)
		if !provinceMap[root] {
			provinceMap[root] = true
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(findCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}))
}
