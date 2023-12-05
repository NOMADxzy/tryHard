package main

import (
	"fmt"
	"strconv"
)

func dfs(nexts [][][]int, visited map[string]bool, pos int, score int) int {
	for _, nextDir := range nexts[pos] {
		if !visited[strconv.Itoa(pos)+"_"+strconv.Itoa(nextDir[0])] {
			visited[strconv.Itoa(pos)+"_"+strconv.Itoa(nextDir[0])] = true
			visited[strconv.Itoa(nextDir[0])+"_"+strconv.Itoa(pos)] = true
			if val := dfs(nexts, visited, nextDir[0], nextDir[1]); val < score {
				score = val
			}
		}
	}
	return score
}

func minScore1(n int, roads [][]int) int {
	nexts := make([][][]int, n)
	for _, road := range roads {
		nexts[road[0]-1] = append(nexts[road[0]-1], []int{road[1] - 1, road[2]})
		nexts[road[1]-1] = append(nexts[road[1]-1], []int{road[0] - 1, road[2]})
	}
	visited := map[string]bool{}
	return dfs(nexts, visited, 0, 1<<31)
}

func minScore(n int, roads [][]int) int {
	parents := make([]int, n)
	score := make([]int, n)
	for i := 0; i < n; i++ {
		score[i] = 1 << 31
		parents[i] = i
	}

	var find func(i int) int
	find = func(i int) int {
		if parents[i] != i {
			parents[i] = find(parents[i])
		}
		return parents[i]
	}
	union := func(i1, i2 int) {
		if score[find(i2)] > score[find(i1)] {
			score[find(i2)] = score[find(i1)]
		}
		parents[find(i1)] = find(i2)
	}

	for _, road := range roads {
		union(road[0]-1, road[1]-1)
		if road[2] < score[find(road[1]-1)] {
			score[find(road[1]-1)] = road[2]
		}
	}
	return score[find(0)]
}

func main() {
	//fmt.Println(minScore1(6, [][]int{{4, 5, 7468}, {6, 2, 7173}, {6, 3, 8365}, {2, 3, 7674}, {5, 6, 7852}, {1, 2, 8547}, {2, 4, 1885}, {2, 5, 5192}, {1, 3, 4065}, {1, 4, 7357}}))
	fmt.Println(minScore(4, [][]int{{1, 2, 2}, {1, 3, 4}, {3, 4, 7}}))
}
