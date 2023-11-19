package main

func findNext(grid, dirs [][]int, x, y int) {
	grid[x][y] = 1
	for _, dir := range dirs {
		nX, nY := x+dir[0], y+dir[1]
		if nX >= 0 && nX < len(grid) && nY >= 0 && nY < len(grid[0]) && grid[nX][nY] == 0 {
			findNext(grid, dirs, nX, nY)
		}
	}
}

func closedIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dirs := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	ans := 0

	for i := 0; i < m; i++ {
		if grid[i][0] == 0 {
			findNext(grid, dirs, i, 0)
		}
		if grid[i][n-1] == 0 {
			findNext(grid, dirs, i, n-1)
		}
	}
	for i := 0; i < n; i++ {
		if grid[0][i] == 0 {
			findNext(grid, dirs, 0, i)
		}
		if grid[m-1][i] == 0 {
			findNext(grid, dirs, m-1, i)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				ans++
				findNext(grid, dirs, i, j)
			}
		}
	}
	return ans
}
