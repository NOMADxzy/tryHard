package main

import "fmt"

func dfs(grid, dirs [][]int, x, y int, hist map[int]bool) bool {
	if hist[x*10+y] {
		return false
	}
	if grid[x][y] == len(grid)*len(grid)-1 {
		return true
	}
	for _, dir := range dirs {
		nX, nY := x+dir[0], y+dir[1]
		if nX >= 0 && nX < len(grid) && nY >= 0 && nY < len(grid[0]) {
			if grid[nX][nY] == grid[x][y] {
				hist[x*10+y] = true
				return false
			} else if grid[nX][nY] == grid[x][y]+1 {
				return dfs(grid, dirs, nX, nY, hist)
			}
		}
	}
	hist[x*10+y] = true
	return false
}

func checkValidGrid(grid [][]int) bool {
	dirs := [][]int{{2, 1}, {2, -1}, {-2, 1}, {-2, -1}, {1, 2}, {1, -2}, {-1, 2}, {-1, -2}}
	if grid[0][0] != 0 {
		return false
	}
	return dfs(grid, dirs, 0, 0, map[int]bool{})
}

func main() {
	grid := [][]int{
		{0, 11, 16, 5, 20},
		{17, 4, 19, 10, 15},
		{12, 1, 8, 21, 6},
		{3, 18, 23, 14, 9},
		{24, 13, 2, 7, 22},
	}
	fmt.Println(checkValidGrid(grid))
}
