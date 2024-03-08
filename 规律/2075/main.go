package main

import "fmt"

func decodeCiphertext(encodedText string, rows int) string {
	if len(encodedText) == 0 {
		return ""
	}
	m, n := rows, len(encodedText)/rows
	grid := make([][]byte, m)
	idx := 0
	for i := 0; i < m; i++ {
		grid[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			grid[i][j] = encodedText[idx]
			idx++
		}
	}
	arr := make([]byte, len(encodedText))
	idx = 0

	for i := 0; i < n; i++ {
		y := i
		for x := 0; x < m && y < n; x++ {
			arr[idx] = grid[x][y]
			idx++
			y++
		}
	}
	var i int
	for i = len(arr) - 1; arr[i] == 0 || arr[i] == ' '; i-- {
	}
	return string(arr[:i+1])
}

func main() {
	fmt.Println(decodeCiphertext("ch   ie   pr", 3))
}
