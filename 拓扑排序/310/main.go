package main

import "fmt"

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	neighbors := make([][]int, n)
	nextCnts := make([]int, n)
	for _, edge := range edges {
		neighbors[edge[0]] = append(neighbors[edge[0]], edge[1])
		neighbors[edge[1]] = append(neighbors[edge[1]], edge[0])
		nextCnts[edge[0]]++
		nextCnts[edge[1]]++
	}
	var queue []int
	for i, cnt := range nextCnts {
		if cnt == 1 {
			queue = append(queue, i)
		}
	}

	moves := 0
	for len(queue) > 0 && moves+len(queue) < n {
		var nextQueue []int
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]
			moves++
			nextCnts[p]--
			for _, q := range neighbors[p] {
				nextCnts[q]--
				if nextCnts[q] == 1 {
					nextQueue = append(nextQueue, q)
				}
			}
		}
		queue = nextQueue
	}
	return queue
}

func main() {
	fmt.Println(findMinHeightTrees(6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}))
}
