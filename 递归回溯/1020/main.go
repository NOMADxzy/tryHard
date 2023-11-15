package main

import "fmt"

func spreadNext(grid [][]int, dirs [][]int, x, y int) {
	grid[x][y] = 0
	for _, dir := range dirs {
		nextX, nextY := x+dir[0], y+dir[1]
		if nextX >= 0 && nextX < len(grid) && nextY >= 0 && nextY < len(grid[0]) && grid[nextX][nextY] == 1 {
			spreadNext(grid, dirs, nextX, nextY)
		}
	}
}

func numEnclaves(grid [][]int) int {
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	ans := 0
	for i := 0; i < len(grid); i++ {
		if grid[i][0] == 1 {
			spreadNext(grid, dirs, i, 0)
		}
		if grid[i][len(grid[0])-1] == 1 {
			spreadNext(grid, dirs, i, len(grid[0])-1)
		}
	}
	for j := 0; j < len(grid[0]); j++ {
		if grid[0][j] == 1 {
			spreadNext(grid, dirs, 0, j)
		}
		if grid[len(grid)-1][j] == 1 {
			spreadNext(grid, dirs, len(grid)-1, j)
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				ans++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(numEnclaves([][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}))
}
