package main

import "fmt"

func findPeakGrid(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	a1, a2 := -1, -1
	var goNext func(i, j int)
	dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	goNext = func(i, j int) {
		if a1 >= 0 {
			return
		}
		bestx, besty := i, j
		for _, dir := range dirs {
			nx, ny := i+dir[0], j+dir[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n && mat[nx][ny] > mat[bestx][besty] {
				bestx, besty = nx, ny
			}
		}
		if bestx == i && besty == j {
			a1, a2 = bestx, besty
		} else {
			goNext(bestx, besty)
		}
	}
	goNext(0, 0)
	return []int{a1, a2}
}

func main() {
	mat := [][]int{{1, 4}, {3, 2}}
	fmt.Println(findPeakGrid(mat))
}
