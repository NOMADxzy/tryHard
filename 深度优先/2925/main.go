package main

import "fmt"

func dfsCnt(nexts [][]int, pos, pre int, childCnts, values []int) int {
	cnt := 0
	for _, p := range nexts[pos] {
		if p != pre {
			cnt += dfsCnt(nexts, p, pos, childCnts, values)
		}
	}
	childCnts[pos] = cnt
	return cnt + values[pos]
}

func dfsVal(nexts [][]int, pos, pre int, childCnts, values []int) int {
	if childCnts[pos] == 0 {
		return 0
	}
	score1 := childCnts[pos]
	score2 := values[pos]
	for _, p := range nexts[pos] {
		if p != pre {
			score2 += dfsVal(nexts, p, pos, childCnts, values)
		}
	}
	return max(score1, score2)
}

func maximumScoreAfterOperations(edges [][]int, values []int) int64 {
	m := len(edges) + 1
	nexts := make([][]int, m)
	for _, edge := range edges {
		nexts[edge[0]] = append(nexts[edge[0]], edge[1])
		nexts[edge[1]] = append(nexts[edge[1]], edge[0])
	}
	childCnts := make([]int, m)
	dfsCnt(nexts, 0, 0, childCnts, values)
	return int64(dfsVal(nexts, 0, 0, childCnts, values))
}

func main() {
	edges := [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 4}, {4, 5}}
	fmt.Println(maximumScoreAfterOperations(edges, []int{5, 2, 5, 2, 1, 1}))
}
