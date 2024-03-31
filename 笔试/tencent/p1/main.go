package main

import (
	"fmt"
)

func solve(n int, edges [][]int) int {
	mark := make([]bool, n)
	for _, edge := range edges {
		mark[edge[0]] = true
		mark[edge[1]] = true
	}
	ans := 0
	for _, b := range mark {
		if !b {
			ans++
		}
	}
	return ans
}

func main() {
	var n, k int
	var p1, p2 int
	var t string
	_, _ = fmt.Scan(&n)
	_, _ = fmt.Scan(&k)
	var edges [][]int
	for i := 0; i < k; i++ {
		_, _ = fmt.Scan(&p1, &p2, &t)
		p1--
		p2--
		if t == "W" {
			edges = append(edges, []int{p1, p2})
		}
	}
	fmt.Println(solve(n, edges))
}
