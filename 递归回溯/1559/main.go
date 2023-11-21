package main

import "fmt"

func goNext(grid [][]byte, idx [][]int, x, y int, id int, dirs [][]int) bool {
	idx[x][y] = id
	for i := 0; i < 4; i++ {
		dir := dirs[i]
		nX, nY := x+dir[0], y+dir[1]
		if nX >= 0 && nX < len(grid) && nY >= 0 && nY < len(grid[0]) && grid[nX][nY] == grid[x][y] {
			if idx[nX][nY] == 0 && goNext(grid, idx, nX, nY, id+1, dirs) {
				return true
			} else if idx[nX][nY] > 0 && id-idx[nX][nY]+1 >= 4 {
				return true
			}
		}
	}
	return false
}

func containsCycle(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])
	dirs := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	idx := make([][]int, m)
	for i := 0; i < m; i++ {
		idx[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if idx[i][j] == 0 && goNext(grid, idx, i, j, 1, dirs) {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(containsCycle([][]byte{{'f', 'a', 'a', 'c', 'b'}, {'e', 'a', 'a', 'e', 'c'}}))
}
