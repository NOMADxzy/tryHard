package main

import "fmt"

// dfs
func shortestPathLength1(graph [][]int) int {
	n := len(graph)
	hist := map[int]int{}
	var dfs func(pos int, state int, acc int)
	ans := 1 << 31
	dfs = func(pos int, state int, acc int) {
		key := state*n + pos
		if state == 1<<n-1 {
			ans = min(ans, acc)
			return
		}
		if v, ok := hist[key]; ok {
			if acc >= v {
				return
			} else {
				hist[key] = acc
			}
		} else {
			hist[key] = acc
		}
		for _, np := range graph[pos] {
			mask := 1 << np
			newState := state
			if state&mask == 0 {
				newState += mask
			}
			dfs(np, newState, acc+1)
		}
	}
	for i := 0; i < n; i++ {
		dfs(i, 1<<i, 0)
	}
	return ans
}

// bfs
func shortestPathLength(graph [][]int) int {
	n := len(graph)
	seen := make([]map[int]bool, n)
	var queue [][]int
	for i := 0; i < n; i++ {
		queue = append(queue, []int{i, 1 << i, 0})
		seen[i] = map[int]bool{}
	}
	endState := 1<<n - 1
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur[1] == endState {
			return cur[2]
		}
		for _, np := range graph[cur[0]] {
			newState := cur[1] | 1<<np
			if !seen[np][newState] {
				queue = append(queue, []int{np, newState, cur[2] + 1})
				seen[np][newState] = true
			}
		}
	}
	return -1
}

func main() {
	graph := [][]int{
		{1, 2, 3},
		{0},
		{0},
		{0},
	}
	fmt.Println(shortestPathLength(graph))
}
