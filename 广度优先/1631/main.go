package main

import (
	"fmt"
	"sort"
)

func minimumEffortPath(heights [][]int) int {
	getAbs := func(a, b int) (ans int) {
		ans = a - b
		if ans < 0 {
			ans = -ans
		}
		return
	}
	m, n := len(heights), len(heights[0])

	useDfs := false

	var dfs func(x, y int, limit int) bool
	var visited map[int]bool
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	dfs = func(x, y int, limit int) bool {
		if x == m-1 && y == n-1 {
			return true
		}
		visited[x*n+y] = true
		arrival := false
		for _, d := range dirs {
			nx, ny := x+d[0], y+d[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n && !visited[nx*n+ny] && getAbs(heights[x][y], heights[nx][ny]) <= limit {
				if dfs(nx, ny, limit) {
					arrival = true
					break
				}
			}
		}
		return arrival
	}
	var bfs func(limit int) bool
	bfs = func(limit int) bool {
		visited = map[int]bool{}
		var queue [][]int
		queue = append(queue, []int{0, 0})
		visited[0] = true
		for len(queue) > 0 {
			var newQueue [][]int
			for len(queue) > 0 {
				point := queue[0]
				queue = queue[1:]
				if point[0] == m-1 && point[1] == n-1 {
					return true
				}
				for _, dir := range dirs {
					nx, ny := point[0]+dir[0], point[1]+dir[1]
					if nx >= 0 && nx < m && ny >= 0 && ny < n && !visited[nx*n+ny] && getAbs(heights[point[0]][point[1]], heights[nx][ny]) <= limit {
						newQueue = append(newQueue, []int{nx, ny})
						visited[nx*n+ny] = true
					}
				}
			}
			queue = newQueue
		}
		return false
	}

	ans := sort.Search(1e6, func(limit int) bool {
		if useDfs {
			visited = map[int]bool{}
			return dfs(0, 0, limit)
		} else {
			return bfs(limit)
		}
	})
	if ans == 1e6 {
		return -1
	}
	return ans
}

func main() {
	heights := [][]int{
		{1, 2, 2},
		{3, 8, 2},
		{5, 3, 5},
	}
	fmt.Println(minimumEffortPath(heights))
}
