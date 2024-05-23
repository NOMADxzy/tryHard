package main

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			x, y := i, j
			pre := matrix[x][y]
			for {
				nx, ny := y, n-1-x
				tmp := matrix[nx][ny]
				matrix[nx][ny] = pre
				pre = tmp
				x, y = nx, ny
				if x == i && y == j {
					break
				}
			}
		}
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
	fmt.Println(matrix)
}
