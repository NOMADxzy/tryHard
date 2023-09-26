package main

import (
	"fmt"
)

func isPacific(heights [][]int, x int, y int) bool {
	if x == 0 || y == 0 {
		return true
	} else {
		return false
	}
}
func isAtlantic(heights [][]int, x int, y int) bool {
	if x == len(heights)-1 || y == len(heights[0])-1 {
		return true
	} else {
		return false
	}
}

func dfs(heights [][]int, x int, y int, history [][]bool, visited [][]bool, dirs [][]int, isWhat func([][]int, int, int) bool) bool {
	if history[x][y] || isWhat(heights, x, y) {
		return true
	}
	visited[x][y] = true

	for i := 0; i < 4; i++ {
		nextX := x + dirs[i][0]
		nextY := y + dirs[i][1]
		if nextX >= 0 && nextX < len(heights) && nextY >= 0 && nextY < len(heights[0]) && heights[x][y] >= heights[nextX][nextY] && !visited[nextX][nextY] {
			if dfs(heights, nextX, nextY, history, visited, dirs, isWhat) {
				history[x][y] = false
				visited[x][y] = false
				return true
			}
		}
	}
	history[x][y] = false
	visited[x][y] = false
	return false
}

func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	history := make([][]bool, m)
	visited := make([][]bool, m)
	for i := 0; i < len(history); i++ {
		history[i] = make([]bool, n)
		visited[i] = make([]bool, n)
	}
	var res [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res1 := dfs(heights, i, j, history, visited, dirs, isPacific)
			res2 := dfs(heights, i, j, history, visited, dirs, isAtlantic)
			if res1 && res2 {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func main() {
	//heights := [][]int{
	//	{1, 2, 2, 3, 5},
	//	{3, 2, 3, 4, 4},
	//	{2, 4, 5, 3, 1},
	//	{6, 7, 1, 4, 5},
	//	{5, 1, 1, 2, 4},
	//}
	//h2 := [][]int{
	//	{3, 3, 3},
	//	{3, 1, 3},
	//	{0, 2, 4},
	//}
	h3 := [][]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}

	fmt.Println(pacificAtlantic(h3))
}
