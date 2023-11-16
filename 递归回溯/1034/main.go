package main

import "fmt"

func findNext(grid, dirs [][]int, target int, x, y int) {
	var valids [][]int
	validsNeighbors := 0
	for _, dir := range dirs {
		nextX, nextY := x+dir[0], y+dir[1]
		if nextX >= 0 && nextX < len(grid) && nextY >= 0 && nextY < len(grid[0]) &&
			(grid[nextX][nextY] == target || grid[nextX][nextY] == 0 || grid[nextX][nextY] == -1) {
			validsNeighbors++
			if grid[nextX][nextY] == target {
				valids = append(valids, []int{nextX, nextY})
			}
		}
	}
	if validsNeighbors == 4 {
		grid[x][y] = 0
	} else {
		grid[x][y] = -1
	}

	for _, valid := range valids {
		findNext(grid, dirs, target, valid[0], valid[1])
	}
}

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	originColor := grid[row][col]
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	m, n := len(grid), len(grid[0])
	findNext(grid, dirs, originColor, row, col)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				grid[i][j] = originColor
			} else if grid[i][j] == -1 {
				grid[i][j] = color
			}
		}
	}
	return grid
}

func main() {
	grid := [][]int{{1, 1}, {1, 2}}
	fmt.Println(colorBorder(grid, 0, 0, 3))
}
