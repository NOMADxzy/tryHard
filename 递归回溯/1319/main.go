package main

import "fmt"

func findNext(nexts map[int][]int, pos int, mark []bool) int {
	if mark[pos] {
		return 0
	}
	cnt := 1
	mark[pos] = true
	for _, p := range nexts[pos] {
		cnt += findNext(nexts, p, mark)
	}
	return cnt
}

func makeConnected(n int, connections [][]int) int {
	nexts := map[int][]int{}
	mark := make([]bool, n)

	for _, connection := range connections {
		nexts[connection[0]] = append(nexts[connection[0]], connection[1])
		nexts[connection[1]] = append(nexts[connection[1]], connection[0])
	}

	totalLines := len(connections)

	needs := 0
	groups := 0
	for i := 0; i < n; i++ {
		val := findNext(nexts, i, mark)
		if val > 0 {
			groups++
			needs += val - 1
		}
	}
	needs += groups - 1
	if needs > totalLines {
		return -1
	}
	return groups - 1
}

func main() {
	fmt.Println(makeConnected(4, [][]int{{0, 1}, {0, 2}, {1, 2}}))
}
