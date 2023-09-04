package main

import "fmt"

func add(queue [][]int, x int, y int, m int, n int, board [][]byte) [][]int {
	if x >= 0 && x < m && y >= 0 && y < n && board[x][y] == 'O' {
		queue = append(queue, []int{x, y})
	}
	return queue
}

func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])

	var queue [][]int

	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			queue = append(queue, []int{i, 0})
		}

		if board[i][n-1] == 'O' {
			queue = append(queue, []int{i, n - 1})
		}
	}

	for j := 1; j < n-1; j++ {
		if board[0][j] == 'O' {
			queue = append(queue, []int{0, j})
		}

		if board[m-1][j] == 'O' {
			queue = append(queue, []int{m - 1, j})
		}
	}

	for len(queue) > 0 {
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]
		board[x][y] = 'M'
		queue = add(queue, x-1, y, m, n, board)
		queue = add(queue, x+1, y, m, n, board)
		queue = add(queue, x, y-1, m, n, board)
		queue = add(queue, x, y+1, m, n, board)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'M' {
				board[i][j] = 'O'
			}
		}
	}

}

func main() {
	//board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	//board := [][]byte{{'O', 'O', 'O'}, {'O', 'O', 'O'}, {'O', 'O', 'O'}}
	board := [][]byte{{'O', 'X', 'X', 'O', 'X'}, {'X', 'O', 'O', 'X', 'O'}, {'X', 'O', 'X', 'O', 'X'}, {'O', 'X', 'O', 'O', 'O'}, {'X', 'X', 'O', 'X', 'O'}}

	solve(board)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			fmt.Print(board[i][j], " ")
		}
		fmt.Println()
	}

}
