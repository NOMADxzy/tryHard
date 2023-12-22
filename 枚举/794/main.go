package main

import "fmt"

func validTicTacToe(board []string) bool {
	Xcnt, Ocnt := 0, 0
	arr := [3][3][2]bool{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 'O' {
				Ocnt++
				arr[i][j][1] = true
			} else if board[i][j] == 'X' {
				Xcnt++
				arr[i][j][0] = true
			}
		}
	}
	var c int
	win := [2]int{}
	for c = 0; c < 2; c++ {
		for i := 0; i < 3; i++ {
			if arr[i][0][c] && arr[i][1][c] && arr[i][2][c] {
				win[c]++
			}
		}
		for j := 0; j < 3; j++ {
			if arr[0][j][c] && arr[1][j][c] && arr[2][j][c] {
				win[c]++
			}
		}
		if arr[0][0][c] && arr[1][1][c] && arr[2][2][c] || arr[0][2][c] && arr[1][1][c] && arr[2][0][c] {
			win[c]++
		}
	}
	if win[0] > 0 && win[1] > 0 {
		return false
	} else if win[0] == 1 {
		return Xcnt == Ocnt+1
	} else if win[1] == 1 {
		return Xcnt == Ocnt
	} else {
		return Xcnt == Ocnt+1 || Xcnt == Ocnt
	}
}

func main() {
	board := []string{"XOX", "OXO", "XOX"}
	fmt.Println(validTicTacToe(board))
}
