package main

import "fmt"

func frogPosition(n int, edges [][]int, t int, target int) float64 {
	nexts := make([][]int, n)
	for _, edge := range edges {
		nexts[edge[0]-1] = append(nexts[edge[0]-1], edge[1]-1)
		nexts[edge[1]-1] = append(nexts[edge[1]-1], edge[0]-1)
	}
	target--
	ans := float64(0)
	var dfs func(pos, step, pre int, p float64)
	dfs = func(pos, step, pre int, p float64) {
		cnt := len(nexts[pos])
		if pre >= 0 {
			cnt--
		}
		if pos == target {
			if step == t || step < t && cnt == 0 {
				ans = p
			}
		} else if step >= t {
			return
		} else {
			for _, np := range nexts[pos] {
				if np != pre {
					dfs(np, step+1, pos, p/float64(cnt))
				}
			}
		}
	}
	dfs(0, 0, -1, 1)
	return ans
}

func main() {
	//edges := [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}}
	edges := [][]int{{2, 1}, {3, 2}, {4, 3}, {5, 3}, {6, 5}, {7, 3}, {8, 4}, {9, 5}}
	fmt.Println(frogPosition(9, edges, 9, 1))
}
