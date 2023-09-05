package main

import "fmt"

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])

	var queue [][]int
	nums := 0

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				nums += 1
				grid[i][j] = '0'
				queue = append(queue, []int{i, j})
				for len(queue) > 0 {
					x, y := queue[0][0], queue[0][1]
					queue = queue[1:]

					for k := 0; k < 4; k++ {
						nx, ny := x+dx[k], y+dy[k]
						if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == '1' {
							grid[nx][ny] = '0'
							queue = append(queue, []int{x + dx[k], y + dy[k]})
						}
					}
				}
			}
		}
	}
	return nums
}

func main() {
	grid := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
	fmt.Println(numIslands(grid))
}
