package main

import "fmt"

func removeStones(stones [][]int) int {
	rowMap := map[int]int{}
	colMap := map[int]int{}

	m := len(stones)
	parents := make([]int, m)
	for i := 0; i < m; i++ {
		parents[i] = i
	}

	var find func(parents []int, index int) int
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

	for i, stone := range stones {
		x, y := stone[0], stone[1]
		if idx, ok := rowMap[y]; ok {
			union(parents, i, idx)
		}
		if idx, ok := colMap[x]; ok {
			union(parents, find(parents, i), idx)
		}
		root := find(parents, i)
		rowMap[y] = root
		colMap[x] = root
	}

	ans := 0
	for i := 0; i < m; i++ {
		if parents[i] == i {
			ans++
		}
	}

	return m - ans
}

func main() {
	fmt.Println(removeStones([][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}))
}
