package main

import "fmt"

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	} else if n == 1 {
		return 1
	}
	getKey := func(x, y int) int {
		return x*100 + y
	}
	dirs := [][]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}
	leftVisited, rightVisited := map[int]bool{}, map[int]bool{}
	var leftQ, rightQ [][]int
	leftVisited[getKey(0, 0)] = true
	rightVisited[getKey(n-1, n-1)] = true
	leftQ = append(leftQ, []int{0, 0})
	rightQ = append(rightQ, []int{n - 1, n - 1})
	step := 1
	for len(leftQ) > 0 && len(rightQ) > 0 {
		step++
		if len(leftQ) <= len(rightQ) {
			var newQ [][]int
			for len(leftQ) > 0 {
				e := leftQ[0]
				leftQ = leftQ[1:]
				for _, dir := range dirs {
					nx, ny := e[0]+dir[0], e[1]+dir[1]
					if nx >= 0 && nx < n && ny >= 0 && ny < n && grid[nx][ny] == 0 && !leftVisited[getKey(nx, ny)] {
						if rightVisited[getKey(nx, ny)] {
							return step
						}
						newQ = append(newQ, []int{nx, ny})
						leftVisited[getKey(nx, ny)] = true
					}
				}
			}
			leftQ = newQ
		} else {
			var newQ [][]int
			for len(rightQ) > 0 {
				e := rightQ[0]
				rightQ = rightQ[1:]
				for _, dir := range dirs {
					nx, ny := e[0]+dir[0], e[1]+dir[1]
					if nx >= 0 && nx < n && ny >= 0 && ny < n && grid[nx][ny] == 0 && !rightVisited[getKey(nx, ny)] {
						if leftVisited[getKey(nx, ny)] {
							return step
						}
						newQ = append(newQ, []int{nx, ny})
						rightVisited[getKey(nx, ny)] = true
					}
				}
			}
			rightQ = newQ
		}
	}
	return -1
}

func main() {
	grid := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}
	fmt.Println(shortestPathBinaryMatrix(grid))
}
