package main

import (
	"fmt"
	"math/bits"
	"sort"
)

func getRectIdx(i, j int) int {
	return 3*(i/3) + j/3
}

type pair struct {
	x        int
	y        int
	possible uint16
}

func bitTostr(board [][]byte) string {
	s := ""
	for i := 0; i < 9; i++ {
		s += string(board[i])
	}
	return s
}

func dfs(board [][]byte, bitset [][]uint16, hist map[string]bool, points []*pair) bool {
	//fmt.Println(remain)
	//state := bitTostr(board)
	//if hist[state] {
	//	return false
	//}
	if len(points) == 0 {
		return true
	}
	sort.Slice(points, func(i, j int) bool {
		return bits.OnesCount16(points[i].possible) < bits.OnesCount16(points[j].possible)
	})
	pa := points[0]
	mask := uint16(1)
	for i := 1; i <= 9; i++ {
		if pa.possible&mask > 0 {
			board[pa.x][pa.y] = byte('0' + i)
			bitset[0][pa.x] &^= mask
			bitset[1][pa.y] &^= mask
			bitset[2][getRectIdx(pa.x, pa.y)] &^= mask
			adjust(points[1:], bitset)
			if dfs(board, bitset, hist, points[1:]) {
				return true
			}
			bitset[0][pa.x] |= mask
			bitset[1][pa.y] |= mask
			bitset[2][getRectIdx(pa.x, pa.y)] |= mask
			board[pa.x][pa.y] = '.'
			adjust(points[1:], bitset)
		}
		mask <<= 1
	}
	//hist[state] = true
	return false
}

func adjust(points []*pair, bitset [][]uint16) {
	for _, point := range points {
		point.possible = bitset[0][point.x] & bitset[1][point.y] & bitset[2][getRectIdx(point.x, point.y)]
	}
}

func solveSudoku(board [][]byte) {
	hist := map[string]bool{}
	bitset := make([][]uint16, 3)
	for i := 0; i < 3; i++ {
		bitset[i] = make([]uint16, 9)
		for j := 0; j < 9; j++ {
			bitset[i][j] = uint16(1)<<9 - 1
		}
	}
	var points []*pair
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				mask := uint16(1) << (board[i][j] - '1')
				bitset[0][i] &^= mask
				bitset[1][j] &^= mask
				bitset[2][getRectIdx(i, j)] &^= mask
			} else {
				points = append(points, &pair{i, j, 0})
			}
		}
	}
	adjust(points, bitset)
	ok := dfs(board, bitset, hist, points)
	fmt.Println(ok)
}

func main() {
	//board := [][]byte{
	//	{'.', '3', '4', '.', '7', '.', '.', '.', '.'},
	//	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	//	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	//	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	//	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	//	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	//	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	//	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	//	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	//}
	board := [][]byte{
		{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
		{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
		{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
		{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
		{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
		{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
		{'.', '.', '.', '2', '7', '5', '9', '.', '.'}}
	solveSudoku(board)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(string(board[i][j]), " ")
		}
		fmt.Println()
	}
}
