package main

import (
	"fmt"
	"strconv"
)

func equalPairs(grid [][]int) int {
	arrMap := map[string]int{}

	same := 0

	for i := 0; i < len(grid); i++ {
		s := ""
		for j := 0; j < len(grid[0]); j++ {
			s += strconv.Itoa(grid[i][j]) + "_"
		}
		arrMap[s] += 1
	}

	for i := 0; i < len(grid[0]); i++ {
		s := ""
		for j := 0; j < len(grid); j++ {
			s += strconv.Itoa(grid[j][i]) + "_"
		}
		same += arrMap[s]
	}

	return same
}

func main() {
	fmt.Println(equalPairs([][]int{
		{3, 1, 2, 2},
		{1, 4, 4, 5},
		{2, 4, 2, 2},
		{2, 4, 2, 2},
	}))
}
