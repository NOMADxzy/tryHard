package main

import "fmt"

func goNext(x, y int, grid [][]int, features [][]bool, dirs [][]int) bool {
	if x == len(grid)-1 && y == len(grid[0])-1 {
		return true
	}
	f := features[grid[x][y]-1]
	grid[x][y] = 0
	for i := 0; i < 4; i++ {
		if f[i] {
			nX, nY := x+dirs[i][0], y+dirs[i][1]
			var i_ int
			if i < 2 {
				i_ = i + 2
			} else {
				i_ = i - 2
			}
			if nX >= 0 && nX < len(grid) && nY >= 0 && nY < len(grid[0]) && grid[nX][nY] > 0 && features[grid[nX][nY]-1][i_] && goNext(nX, nY, grid, features, dirs) {
				return true
			}
		}
	}
	return false
}

func hasValidPath(grid [][]int) bool {
	features := [][]bool{
		{false, true, false, true},
		{true, false, true, false},
		{false, false, true, true},
		{false, true, true, false},
		{true, false, false, true},
		{true, true, false, false},
	}
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	return goNext(0, 0, grid, features, dirs)
}

func main() {
	grid := [][]int{{2, 4, 3}, {6, 5, 2}}
	fmt.Println(hasValidPath(grid))
}
