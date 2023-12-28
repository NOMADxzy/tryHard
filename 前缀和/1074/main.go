package main

import "fmt"

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	m, n := len(matrix), len(matrix[0])
	vsum := make([][]int, m)
	for i := 0; i < m; i++ {
		vsum[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				vsum[i][j] = matrix[i][j]
			} else {
				vsum[i][j] = vsum[i-1][j] + matrix[i][j]
			}
		}
	}
	ans := 0
	for top := 0; top < m; top++ {
		for bottom := top; bottom < m; bottom++ {
			vMap := map[int]int{}
			acc := 0
			for j := 0; j < n; j++ {
				acc += vsum[bottom][j]
				if top > 0 {
					acc -= vsum[top-1][j]
				}
				if vMap[acc-target] > 0 {
					ans += vMap[acc-target]
				}
				vMap[acc]++
				if acc == target {
					ans++
				}
			}
		}
	}
	return ans
}

func main() {
	mat := [][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}
	fmt.Println(numSubmatrixSumTarget(mat, 0))
}
