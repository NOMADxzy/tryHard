package main

import (
	"fmt"
)

func solve(n int, edges [][]int) int {
	nexts := make([][]int, n)
	for _, edge := range edges {
		f, t := edge[0], edge[1]
		nexts[f] = append(nexts[f], t)
		nexts[t] = append(nexts[t], f)
	}
	mark := make([]bool, n)
	var dfs func(pos int) int
	dfs = func(pos int) int {
		cnt := 1
		mark[pos] = true
		for _, np := range nexts[pos] {
			if !mark[np] {
				cnt += dfs(np)
			}
		}
		return cnt
	}
	var groups []int
	for i := 0; i < n; i++ {
		if !mark[i] {
			cnt := dfs(i)
			groups = append(groups, cnt)
			if len(groups) > 2 {
				return 0
			}
		}
	}
	return groups[0] * groups[1]
}

func main() {
	var n, m int
	var p1, p2 int
	_, _ = fmt.Scan(&n, &m)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		_, _ = fmt.Scan(&p1, &p2)
		p1--
		p2--
		edges[i] = []int{p1, p2}
	}
	fmt.Println(solve(n, edges))
}
