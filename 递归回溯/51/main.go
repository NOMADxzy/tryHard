package main

import "fmt"

func findNext(answers *[][]string, history []int, pos int, n int, model []byte) {
	if pos == n {
		answer := make([]string, n)
		for i := 0; i < n; i++ {
			model[history[i]] = 'Q'
			answer[i] = string(model)
			model[history[i]] = '.'
		}
		*answers = append(*answers, answer)
		return
	}
	for i := 0; i < n; i++ {
		valid := true
		for j := 0; j < pos; j++ {
			if history[j] == i || history[j]-i == j-pos || history[j]-i == pos-j {
				valid = false
				break
			}
		}
		if valid {
			history[pos] = i
			findNext(answers, history, pos+1, n, model)
		}
	}
}

func solveNQueens(n int) [][]string {
	history := make([]int, n)

	var ans [][]string
	model := make([]byte, n)
	for i := 0; i < n; i++ {
		model[i] = '.'
	}
	findNext(&ans, history, 0, n, model)

	return ans
}

func main() {
	n := 4
	board := solveNQueens(n)
	for i := 0; i < len(board); i++ {
		fmt.Println(board[i])
	}
}
