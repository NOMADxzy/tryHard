package main

import "fmt"

func regionsBySlashes(grid []string) int {
	n := len(grid)
	mat := make([][][4]bool, n)
	for i := 0; i < n; i++ {
		mat[i] = make([][4]bool, n)
	}
	ans := 0

	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	//dirs := map[byte][]int{'/',[]int}
	nexts := map[int]int{0: 2, 2: 0, 1: 3, 3: 1}
	neighbors1 := map[int]int{0: 3, 3: 0, 1: 2, 2: 1}
	neighbors2 := map[int]int{0: 1, 1: 0, 3: 2, 2: 3}

	var dfs func(x, y, z int)
	dfs = func(x, y int, z int) {
		mat[x][y][z] = true

		dir := dirs[z]
		nx, ny := x+dir[0], y+dir[1]
		if nx >= 0 && nx < n && ny >= 0 && ny < n && !mat[nx][ny][nexts[z]] {
			dfs(nx, ny, nexts[z])
		}
		if grid[x][y] == '/' {
			if !mat[x][y][neighbors1[z]] {
				dfs(x, y, neighbors1[z])
			}
		} else if grid[x][y] == '\\' {
			if !mat[x][y][neighbors2[z]] {
				dfs(x, y, neighbors2[z])
			}
		} else {
			for nz := 0; nz < 4; nz++ {
				if !mat[x][y][nz] {
					dfs(x, y, nz)
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < 4; k++ {
				if !mat[i][j][k] {
					ans++
					dfs(i, j, k)
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(regionsBySlashes([]string{"/\\", "\\/"}))
}
