package main

import "fmt"

func cntNext(nexts [][]int, pos int, visited []bool) int {
	cnt := 1
	visited[pos] = true
	for _, np := range nexts[pos] {
		if !visited[np] {
			cnt += cntNext(nexts, np, visited)
		}
	}
	return cnt
}

func reachableNodes(n int, edges [][]int, restricted []int) int {
	nexts := make([][]int, n)
	visited := make([]bool, n)
	for _, edge := range edges {
		nexts[edge[0]] = append(nexts[edge[0]], edge[1])
		nexts[edge[1]] = append(nexts[edge[1]], edge[0])
	}
	for _, p := range restricted {
		visited[p] = true
	}
	if visited[0] {
		return 0
	}
	return cntNext(nexts, 0, visited)
}

func main() {
	fmt.Println(reachableNodes(7, [][]int{{0, 1}, {1, 2}, {3, 1}, {4, 0}, {0, 5}, {5, 6}}, []int{4, 5}))
}
