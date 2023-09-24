package main

import "fmt"

func countAround(board [][]int, i int, j int, di []int, dj []int) int {
	m, n := len(board), len(board[0])
	cnt := 0
	for k := 0; k < 8; k++ {
		nextI, nextJ := i+di[k], j+dj[k]
		if nextI >= 0 && nextI < m && nextJ >= 0 && nextJ < n && (board[nextI][nextJ] == 1 || board[nextI][nextJ] == -1) {
			cnt++
		}
	}
	return cnt
}

func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	di := []int{-1, -1, -1, 0, 1, 1, 1, 0}
	dj := []int{1, 0, -1, -1, -1, 0, 1, 1}

	//0,-1, 1,2

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cnt := countAround(board, i, j, di, dj)
			if board[i][j] == 0 && cnt == 3 {
				board[i][j] = 2
			}
			if board[i][j] == 1 {
				if cnt < 2 || cnt > 3 {
					board[i][j] = -1
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == -1 {
				board[i][j] = 0
			} else if board[i][j] == 2 {
				board[i][j] = 1
			}
		}
	}
}

func main() {
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0}}
	gameOfLife(board)
	fmt.Println("done")
}
