package main

import "fmt"

func goNext(childs [][][]int, mark []bool, pos int) int {
	mark[pos] = true
	cnt := 0
	for _, pair := range childs[pos] {
		nextP, invalid := pair[0], pair[1]
		if !mark[nextP] {
			cnt += invalid + goNext(childs, mark, nextP)
		}
	}
	return cnt
}

func minReorder(n int, connections [][]int) int {
	childs := make([][][]int, n)
	mark := make([]bool, n)
	for _, connection := range connections {
		f, t := connection[0], connection[1]
		childs[f] = append(childs[f], []int{t, 1})
		childs[t] = append(childs[t], []int{f, 0})
	}
	return goNext(childs, mark, 0)
}

func main() {
	fmt.Println(minReorder(6, [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}))
}
