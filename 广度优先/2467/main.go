package main

import "fmt"

func findDir(pos int, nexts [][]int, marked []bool) []int {
	if pos == 0 {
		return []int{0}
	}
	marked[pos] = true
	for _, p := range nexts[pos] {
		if !marked[p] {
			if path := findDir(p, nexts, marked); path != nil {
				return append(path, pos)
			}
		}
	}
	return nil
}

func dfs(nexts [][]int, visited []bool, rightMap map[int]int, pos, curStep int, amount []int) int {
	visited[pos] = true
	Score := 0
	if rstep, ok := rightMap[pos]; ok && rstep == curStep {
		Score += amount[pos] / 2
	} else if !ok || rstep > curStep {
		Score += amount[pos]
	}
	bestChildSocre := -1 << 31
	for _, p := range nexts[pos] {
		if !visited[p] {
			if bestScore := dfs(nexts, visited, rightMap, p, curStep+1, amount); bestScore > bestChildSocre {
				bestChildSocre = bestScore
			}
		}
	}
	if bestChildSocre == -1<<31 {
		bestChildSocre = 0
	}
	return bestChildSocre + Score
}

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	n := len(edges) + 1
	nexts := make([][]int, n)
	for _, edge := range edges {
		nexts[edge[0]] = append(nexts[edge[0]], edge[1])
		nexts[edge[1]] = append(nexts[edge[1]], edge[0])
	}
	leftVisited, rightVisited := make([]int, n), make([]int, n)
	leftVisited[0] = 1
	rightVisited[bob] = 1
	rightPaths := findDir(bob, nexts, make([]bool, n))
	for i := 0; i < len(rightPaths)/2; i++ {
		rightPaths[i], rightPaths[len(rightPaths)-1-i] = rightPaths[len(rightPaths)-1-i], rightPaths[i]
	}
	rightMap := map[int]int{}
	for i, p := range rightPaths {
		rightMap[p] = i
	}

	return dfs(nexts, make([]bool, n), rightMap, 0, 0, amount)
}

func main() {
	edges := [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}
	fmt.Println(mostProfitablePath(edges, 3, []int{-2, 4, 2, -4, 6}))
}
