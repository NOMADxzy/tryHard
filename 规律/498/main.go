package main

import "fmt"

func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	ans := make([]int, m*n)
	dir := []int{-1, 1}
	x, y := 0, 0
	for idx := 0; idx < len(ans); idx++ {
		ans[idx] = mat[x][y]
		xIsBorder := x == 0 || x == m-1
		yIsBorder := y == 0 || y == n-1
		for x+dir[0] < 0 || x+dir[0] >= m || y+dir[1] < 0 || y+dir[1] >= n {
			dir[0], dir[1] = -dir[0], -dir[1]
			if xIsBorder && yIsBorder {
				if x == m-1 {
					y++
				} else if y == n-1 {
					x++
				} else {
					y++
				}
			} else if xIsBorder {
				y++
			} else {
				x++
			}
			if x >= m || y >= n {
				break
			}
			idx++
			ans[idx] = mat[x][y]
			xIsBorder = x == 0 || x == m-1
			yIsBorder = y == 0 || y == n-1
		}
		x, y = x+dir[0], y+dir[1]
	}
	return ans
}

func main() {
	fmt.Println(findDiagonalOrder([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}))
}
