package main

import "fmt"

func maximalRectangle(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	mat := make([][]int, m)
	for i := 0; i < m; i++ {
		mat[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				mat[i][j] = 1
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			if mat[i][j] == 1 {
				mat[i][j] = mat[i][j-1] + 1
			}
		}
	}
	ans := 0
	for j := 0; j < n; j++ {
		var incStack []int
		nextSmaller := make([]int, m)
		for i := 0; i < m; i++ {
			for len(incStack) > 0 && mat[incStack[len(incStack)-1]][j] > mat[i][j] {
				idx := incStack[len(incStack)-1]
				incStack = incStack[:len(incStack)-1]
				nextSmaller[idx] = i
			}
			nextSmaller[i] = -1
			incStack = append(incStack, i)
		}
		var decStack []int
		preSmaller := make([]int, m)
		for i := m - 1; i >= 0; i-- {
			for len(decStack) > 0 && mat[decStack[len(decStack)-1]][j] > mat[i][j] {
				idx := decStack[len(decStack)-1]
				decStack = decStack[:len(decStack)-1]
				preSmaller[idx] = i
			}
			preSmaller[i] = -1
			decStack = append(decStack, i)
		}
		for i := 0; i < m; i++ {
			right := m
			left := -1
			if nextSmaller[i] >= 0 {
				right = nextSmaller[i]
			}
			if preSmaller[i] >= 0 {
				left = preSmaller[i]
			}
			if val := (right - left - 1) * mat[i][j]; val > ans {
				ans = val
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maximalRectangle([][]byte{
		{'1', '0', '1', '1', '0', '1'},
		{'1', '1', '1', '1', '1', '1'},
		{'0', '1', '1', '0', '1', '1'},
		{'1', '1', '1', '0', '1', '0'},
		{'0', '1', '1', '1', '1', '1'},
		{'1', '1', '0', '1', '1', '1'},
	}))
}
