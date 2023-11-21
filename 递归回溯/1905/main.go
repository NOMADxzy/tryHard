package main

import "fmt"

func findNext(grid1, grid2, dirs [][]int, x, y int) bool {
	valid := true
	if grid1[x][y] == 0 {
		valid = false
	}
	grid2[x][y] = 0
	for _, dir := range dirs {
		nX, nY := x+dir[0], y+dir[1]
		if nX >= 0 && nX < len(grid1) && nY >= 0 && nY < len(grid1[0]) && grid2[nX][nY] == 1 && !findNext(grid1, grid2, dirs, nX, nY) {
			valid = false
		}
	}
	return valid
}

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	dirs := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	ans := 0
	for i := 0; i < len(grid1); i++ {
		for j := 0; j < len(grid1[0]); j++ {
			if grid2[i][j] == 1 && findNext(grid1, grid2, dirs, i, j) {
				ans++
			}
		}
	}
	return ans
}

func main() {
	grid1 := [][]int{{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}}
	grid2 := [][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0}}

	fmt.Println(countSubIslands(grid1, grid2))
}
