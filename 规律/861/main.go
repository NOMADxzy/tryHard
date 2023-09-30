package main

import "fmt"

func matrixScore(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	moveOfEachRow := make([]bool, m)
	//moveOfEachRow2 := make([]bool, m)
	layer := 1
	for i := 1; i < n; i++ {
		layer *= 2
	}
	score := 0

	for i := 0; i < m; i++ {
		moveOfEachRow[i] = grid[i][0] == 0
		score += layer
	}

	for j := 1; j < n; j++ {
		layer /= 2
		count1 := 0
		for i := 0; i < m; i++ {
			if moveOfEachRow[i] && grid[i][j] == 0 || !moveOfEachRow[i] && grid[i][j] == 1 {
				count1++
			}
		}
		if count1 <= m/2 {
			count1 = m - count1
		}
		score += layer * count1
	}
	return score
}

func main() {
	grid := [][]int{
		{0, 0, 1, 1},
		{1, 0, 1, 0},
		{1, 1, 0, 0},
	}
	fmt.Println(matrixScore(grid))
}
