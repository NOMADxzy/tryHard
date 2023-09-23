package main

import "fmt"

func dfs(matrix [][]int, i int, j int, cache [][]int, di []int, dj []int) int {
	if cache[i][j] > 0 {
		return cache[i][j]
	}
	m, n := len(matrix), len(matrix[0])
	maxLen := 0
	for k := 0; k < 4; k++ {
		nextX, nextY := i+di[k], j+dj[k]
		if nextX < m && nextX >= 0 && nextY < n && nextY >= 0 && matrix[i][j] < matrix[nextX][nextY] {
			length := dfs(matrix, nextX, nextY, cache, di, dj)
			if length > maxLen {
				maxLen = length
			}
		}
	}

	cache[i][j] = maxLen + 1
	return maxLen + 1
}

func longestIncreasingPath(matrix [][]int) int {
	di := []int{-1, 0, 1, 0}
	dj := []int{0, 1, 0, -1}

	m, n := len(matrix), len(matrix[0])
	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
	}

	best := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			l := dfs(matrix, i, j, cache, di, dj)
			if l > best {
				best = l
			}
		}
	}
	return best
}

func main() {
	fmt.Println(longestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}))
}
