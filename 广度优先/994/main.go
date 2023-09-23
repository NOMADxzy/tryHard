package main

import "fmt"

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var curQueue [][]int
	var nextQueue [][]int

	dm := []int{-1, 0, 1, 0}
	dn := []int{0, 1, 0, -1}
	time := -1
	r0 := true
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				curQueue = append(curQueue, []int{i, j})
			} else if grid[i][j] == 1 {
				r0 = false
			}
		}
	}
	if r0 {
		return 0
	}

	for len(curQueue) > 0 {
		for len(curQueue) > 0 {
			pos := curQueue[0]
			curQueue = curQueue[1:]
			for i := 0; i < 4; i++ {
				neighX, neighY := pos[0]+dm[i], pos[1]+dn[i]
				if neighX >= 0 && neighX < m && neighY >= 0 && neighY < n && grid[neighX][neighY] == 1 {
					grid[neighX][neighY] = 2
					nextQueue = append(nextQueue, []int{neighX, neighY})
				}
			}
		}
		curQueue = nextQueue
		nextQueue = nil
		time++
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return time
}

func main() {
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}))
}
