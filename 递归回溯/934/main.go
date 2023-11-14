package main

import "fmt"

func changeTo2(grid [][]int, x int, y int, dirs [][]int) {
	grid[x][y] = 2
	for _, dir := range dirs {
		nextX, nextY := x+dir[0], y+dir[1]
		if nextX >= 0 && nextX < len(grid) && nextY >= 0 && nextY < len(grid[0]) && grid[nextX][nextY] == 1 {
			changeTo2(grid, nextX, nextY, dirs)
		}
	}
}

func shortestBridge(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	var queue [][]int
outLayer:
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				changeTo2(grid, i, j, dirs)
				break outLayer
			}
		}
	}
	var epoch int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	var nextQ [][]int
mainLoop:
	for epoch = 0; ; epoch++ {
		nextQ = nil
		for len(queue) > 0 {
			e := queue[0]
			queue = queue[1:]
			for _, dir := range dirs {
				nextX, nextY := e[0]+dir[0], e[1]+dir[1]
				if nextX >= 0 && nextX < len(grid) && nextY >= 0 && nextY < len(grid[0]) {
					if grid[nextX][nextY] == 0 {
						grid[nextX][nextY] = 2
						nextQ = append(nextQ, []int{nextX, nextY})
					} else if grid[nextX][nextY] == 1 {
						break mainLoop
					}
				}
			}
		}
		queue = nextQ
	}
	return epoch
}

func main() {
	grid := [][]int{{1, 0}, {0, 1}}
	fmt.Println(shortestBridge(grid))
}
