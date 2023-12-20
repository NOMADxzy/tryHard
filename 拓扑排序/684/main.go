package main

import "fmt"

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	degrees := make([]int, n+1)
	nexts := make([][]int, n+1)
	for _, edge := range edges {
		nexts[edge[0]] = append(nexts[edge[0]], edge[1])
		degrees[edge[0]]++
		nexts[edge[1]] = append(nexts[edge[1]], edge[0])
		degrees[edge[1]]++
	}
	deleted := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		deleted[i] = make([]bool, n+1)
	}
	var queue []int
	for i := 1; i <= n; i++ {
		if degrees[i] == 1 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		var newQueue []int
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]
			for _, np := range nexts[p] {
				degrees[p]--
				degrees[np]--
				deleted[p][np] = true
				if degrees[np] == 1 {
					newQueue = append(newQueue, np)
				}
			}
		}
		queue = newQueue
	}
	for i := n - 1; i >= 0; i-- {
		edge := edges[i]
		if !(deleted[edge[0]][edge[1]] || deleted[edge[1]][edge[0]]) {
			return edge
		}
	}
	return []int{-1, -1}
}

func main() {
	edges := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}
	fmt.Println(findRedundantConnection(edges))
}
