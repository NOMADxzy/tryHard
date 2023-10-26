package main

import "fmt"

func maxDistance(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	//marked := make([][]bool, m)
	//for i := 0; i < m; i++ {
	//	marked[i] = make([]bool, n)
	//}

	dir := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for dist := 1; ; dist++ {
		flag := false
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] == dist {
					//marked[i][j] = true
					for d := 0; d < 4; d++ {
						nextX, nextY := i+dir[d][0], j+dir[d][1]
						if nextX >= 0 && nextX < m && nextY >= 0 && nextY < n && grid[nextX][nextY] == 0 {
							flag = true
							grid[nextX][nextY] = dist + 1
						}
					}
				}
			}
		}
		if !flag {
			if dist == 1 {
				return -1
			} else {
				return dist - 1
			}
		}
	}
}

func main() {
	grid := [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}
	fmt.Println(maxDistance(grid))
}
