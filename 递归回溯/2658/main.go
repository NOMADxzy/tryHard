package main

import "fmt"

func findMaxFish(grid [][]int) int {
	var ans int
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	var sum func(x, y int) int
	m, n := len(grid), len(grid[0])

	sum = func(x, y int) int {
		cnt := grid[x][y]
		grid[x][y] = 0
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] > 0 {
				cnt += sum(nx, ny)
			}
		}
		return cnt
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				ans = max(ans, sum(i, j))
			}
		}
	}
	return ans
}

func main() {
	grid := [][]int{
		{0, 2, 1, 0},
		{4, 0, 0, 3},
		{1, 0, 0, 4},
		{0, 3, 2, 0},
	}
	fmt.Println(findMaxFish(grid))
}
