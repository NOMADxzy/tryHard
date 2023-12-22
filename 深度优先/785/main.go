package main

import "fmt"

func checkNext(graph [][]int, pos int, leftMark, rightMark map[int]bool, curLeft bool) bool {
	if curLeft && rightMark[pos] || !curLeft && leftMark[pos] {
		return false
	} else if leftMark[pos] || rightMark[pos] {
		return true
	}
	if curLeft {
		leftMark[pos] = true
	} else {
		rightMark[pos] = true
	}
	for _, np := range graph[pos] {
		if !checkNext(graph, np, leftMark, rightMark, !curLeft) {
			return false
		}
	}
	return true
}

func isBipartite(graph [][]int) bool {
	n := len(graph)
	leftMark := map[int]bool{}
	rightMark := map[int]bool{}
	for i := 0; i < n; i++ {
		if (!leftMark[i] && !rightMark[i]) && !checkNext(graph, i, leftMark, rightMark, true) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isBipartite([][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}))
}
