package main

import "fmt"

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	rnexts := make([][]int, n)
	bnexts := make([][]int, n)
	for i := 0; i < n; i++ {
		rnexts[i] = make([]int, n)
		bnexts[i] = make([]int, n)
	}
	for _, edge := range redEdges {
		f, t := edge[0], edge[1]
		rnexts[f][t]++
	}
	for _, edge := range blueEdges {
		f, t := edge[0], edge[1]
		bnexts[f][t]++
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}
	bfs := func(red bool) []int {
		var queue []int
		queue = append(queue, 0)
		step := 0
		for len(queue) > 0 {
			var newQ []int
			for len(queue) > 0 {
				e := queue[0]
				if ans[e] < 0 || step < ans[e] {
					ans[e] = step
				}
				queue = queue[1:]
				if red {
					for j := 0; j < n; j++ {
						if rnexts[e][j] > 0 {
							newQ = append(newQ, j)
							rnexts[e][j]--
						}
					}
				} else {
					for j := 0; j < n; j++ {
						if bnexts[e][j] > 0 {
							newQ = append(newQ, j)
							bnexts[e][j]--
						}
					}
				}
			}
			red = !red
			queue = newQ
			step++
		}
		return ans
	}
	bfs(true)
	for i := 0; i < n; i++ {
		clear(rnexts[i])
		clear(bnexts[i])
	}
	for _, edge := range redEdges {
		f, t := edge[0], edge[1]
		rnexts[f][t]++
	}
	for _, edge := range blueEdges {
		f, t := edge[0], edge[1]
		bnexts[f][t]++
	}
	bfs(false)
	return ans
}

func main() {
	fmt.Println(shortestAlternatingPaths(3, [][]int{{0, 1}, {1, 2}}, [][]int{}))
}
