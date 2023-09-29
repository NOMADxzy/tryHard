package main

import "fmt"

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	ans := make([][]int, m)
	var curLayer, nextLayer [][]int
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				curLayer = append(curLayer, []int{i, j})
			} else {
				ans[i][j] = -1
			}
		}
	}
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	dist := 1
	for len(curLayer) > 0 {
		for len(curLayer) > 0 {
			x, y := curLayer[0][0], curLayer[0][1]
			curLayer = curLayer[1:]
			for i := 0; i < 4; i++ {
				nextX, nextY := x+dirs[i][0], y+dirs[i][1]
				if nextX >= 0 && nextX < m && nextY >= 0 && nextY < n && mat[nextX][nextY] == 1 && ans[nextX][nextY] == -1 {
					ans[nextX][nextY] = dist
					nextLayer = append(nextLayer, []int{nextX, nextY})
				}
			}
		}
		curLayer = nextLayer
		nextLayer = nil
		dist++
	}
	return ans
}

func main() {
	mat := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}
	fmt.Println(updateMatrix(mat))
}
