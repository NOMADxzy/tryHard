package main

import "fmt"

func isRobotBounded(instructions string) bool {
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	changeDir := func(dirIdx int, left bool) int {
		if left {
			dirIdx -= 1
			if dirIdx < 0 {
				dirIdx = 3
			}
		} else {
			dirIdx += 1
			if dirIdx > 3 {
				dirIdx = 0
			}
		}
		return dirIdx
	}
	dirIdx := 0
	x, y := 0, 0
	for i := 0; i < len(instructions); i++ {
		if instructions[i] == 'G' {
			x, y = x+dirs[dirIdx][0], y+dirs[dirIdx][1]
		} else {
			dirIdx = changeDir(dirIdx, instructions[i] == 'L')
		}
	}
	return x == 0 && y == 0 || dirIdx != 0
}

func main() {
	fmt.Println(isRobotBounded("GGLLGG"))
}
