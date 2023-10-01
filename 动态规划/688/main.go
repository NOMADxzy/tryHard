package main

import "fmt"

func cptProbability(probMat [][]float64, n, x, y int, dirs [][]int) float64 {
	p := float64(0)
	for i := 0; i < 8; i++ {
		nX, nY := x+dirs[i][0], y+dirs[i][1]
		if nX >= 0 && nX < n && nY >= 0 && nY < n {
			p += probMat[nX][nY]
		}
	}
	return p / float64(8)
}

func knightProbability(n int, k int, row int, column int) float64 {
	dirs := [][]int{{-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}}
	probMat := make([][]float64, n)
	newProbMat := make([][]float64, n)
	for i := 0; i < n; i++ {
		probMat[i] = make([]float64, n)
		newProbMat[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			probMat[i][j] = 1
		}
	}
	for epoch := k - 1; epoch >= 0; epoch-- {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				newProbMat[i][j] = cptProbability(probMat, n, i, j, dirs)
			}
		}
		probMat = newProbMat
		newProbMat = make([][]float64, n)
		for j := 0; j < n; j++ {
			newProbMat[j] = make([]float64, n)
		}
	}
	return probMat[row][column]
}

func main() {
	fmt.Println(knightProbability(3, 2, 0, 0))
}
