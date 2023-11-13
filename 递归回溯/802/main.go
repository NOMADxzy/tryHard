package main

import "fmt"

func findNext(mark []int, visited []bool, graph [][]int, pos int) bool {
	if mark[pos] != 0 {
		return mark[pos] > 0
	}
	if visited[pos] {
		return false
	}
	if len(graph[pos]) > 0 {
		for _, nextP := range graph[pos] {
			visited[pos] = true
			if !findNext(mark, visited, graph, nextP) {
				mark[pos] = -1
				visited[pos] = false
				return false
			}
			visited[pos] = false
		}
	}
	mark[pos] = 1
	return true
}

func eventualSafeNodes(graph [][]int) []int {
	m := len(graph)
	mark := make([]int, len(graph))
	var ans []int
	visited := make([]bool, m)
	for i := 0; i < m; i++ {
		if findNext(mark, visited, graph, i) {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	graph := [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}
	fmt.Println(eventualSafeNodes(graph))
}
