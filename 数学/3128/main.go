package main

import "fmt"

func numberOfRightTriangles(grid [][]int) int64 {
	m, n := len(grid), len(grid[0])
	rowCnt1, colCnt1 := make([]int, m), make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				rowCnt1[i]++
				colCnt1[j]++
			}
		}
	}
	ans := int64(0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				ans += int64(rowCnt1[i]-1) * int64(colCnt1[j]-1)
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfRightTriangles([][]int{{1, 0, 1}, {1, 0, 0}, {1, 0, 0}}))
}
