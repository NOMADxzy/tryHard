package main

import "fmt"

func dfsCountNext(grid [][]int, x, y int, dirs [][]int) int {
	cnt := 1
	grid[x][y] = 0
	for i := 0; i < 4; i++ {
		nextX, nextY := x+dirs[i][0], y+dirs[i][1]
		if nextX >= 0 && nextX < len(grid) && nextY >= 0 && nextY < len(grid[0]) && grid[nextX][nextY] == 1 {
			cnt += dfsCountNext(grid, nextX, nextY, dirs)
		}
	}
	return cnt
}

func maxAreaOfIsland(grid [][]int) int {
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				area := dfsCountNext(grid, i, j, dirs)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

func main() {
	grid := [][]int{
		{1, 1, 0, 0},
		{1, 0, 0, 0},
		{0, 0, 1, 1},
		{0, 0, 1, 1},
	}
	fmt.Println(maxAreaOfIsland(grid))
}
