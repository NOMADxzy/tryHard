package main

import "fmt"

func largestArea(grid []string) int {
	var dfs func(x, y int) (int, bool)
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	dfs = func(x, y int) (int, bool) {
		cnt := 1
		valid := true
		cur := grid[x][y]
		visited[x][y] = true
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if nx < 0 || nx >= len(grid) || ny < 0 || ny >= len(grid[0]) || grid[nx][ny] == '0' {
				valid = false
			} else if !visited[nx][ny] {
				if grid[nx][ny] == cur {
					c, b := dfs(nx, ny)
					if !b {
						valid = false
						cnt += c
					}
					cnt += c
				}

			}
		}
		return cnt, valid
	}
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !visited[i][j] && grid[i][j] != '0' {
				cnt, valid := dfs(i, j)
				if valid && cnt > ans {
					ans = cnt
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(largestArea([]string{"11111100000", "21243101111", "21224101221", "11111101111"}))
}
