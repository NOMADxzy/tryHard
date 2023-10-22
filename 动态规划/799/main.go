package main

import "fmt"

func champagneTower(poured int, query_row int, query_glass int) float64 {
	tower := make([][]float64, query_row+1)
	for i := 0; i <= query_row; i++ {
		tower[i] = make([]float64, i+1)
	}

	tower[0][0] = float64(poured)
	for i := 0; i < query_row; i++ {
		for j := 0; j <= i; j++ {
			if tower[i][j] > 1 {
				tower[i+1][j] += (tower[i][j] - 1) / 2
				tower[i+1][j+1] += (tower[i][j] - 1) / 2
			}
		}
	}
	if tower[query_row][query_glass] > 1 {
		return 1
	} else {
		return tower[query_row][query_glass]
	}
}

func main() {
	fmt.Println(champagneTower(100000009, 33, 17))
}
