package main

import "fmt"

func dfs(board [][]byte, x int, y int, dirs [][]int) {
	board[x][y] = '.'
	for i := 0; i < 4; i++ {
		nextX, nextY := x+dirs[i][0], y+dirs[i][1]
		if nextX >= 0 && nextX < len(board) && nextY >= 0 && nextY < len(board[0]) && board[nextX][nextY] == 'X' {
			dfs(board, nextX, nextY, dirs)
		}
	}
}

func countBattleships(board [][]byte) int {
	cnt := 0
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'X' {
				cnt++
				dfs(board, i, j, dirs)
			}
		}
	}
	return cnt
}

func main() {
	board := [][]byte{
		{'X', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
	}
	fmt.Println(countBattleships(board))
}
