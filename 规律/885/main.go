package main

import "fmt"

func nextDir(dir []int) int {
	x, y := dir[0], dir[1]
	step := 0
	if x == 0 && y > 0 {
		dir[0] = y
		dir[1] = 0
		step = 1
	} else if x > 0 && y == 0 {
		dir[0] = 0
		dir[1] = -(x + 1)
		step = -1
	} else if x == 0 && y < 0 {
		dir[0] = y
		dir[1] = 0
		step = -1
	} else {
		dir[0] = 0
		dir[1] = -x + 1
		step = 1
	}
	return step
}

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	dir := []int{0, 1}
	cnt := rows * cols
	ans := make([][]int, cnt)
	step := 1

	x, y := rStart, cStart
	for i := 0; i < cnt; {
		if dir[0] == 0 {
			for nextY := y + dir[1]; y != nextY; y += step {
				if x >= 0 && x < rows && y >= 0 && y < cols {
					ans[i] = []int{x, y}
					i++
				}
			}
		} else {
			for nextX := x + dir[0]; x != nextX; x += step {
				if x >= 0 && x < rows && y >= 0 && y < cols {
					ans[i] = []int{x, y}
					i++
				}
			}
		}
		step = nextDir(dir)
	}
	return ans
}

func main() {
	fmt.Println(spiralMatrixIII(5, 6, 1, 4))
}
