package main

import "fmt"

func isValid(x int, y int, maze [][]byte) bool {
	if x >= 0 && x < len(maze) && y >= 0 && y < len(maze[0]) && maze[x][y] == '.' {
		return true
	}
	return false
}

func getScores(x int, y int, maze [][]byte) int {
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}
	cnt := 0
	for i := 0; i < 4; i++ {
		if isValid(x+dx[i], y+dy[i], maze) {
			cnt++
		}
	}
	return cnt
}

func goNext(maze [][]byte, pre int, x int, y int) {

	if x != 0 && x != len(maze)-1 && y != 0 && y != len(maze[0])-1 {

	} else {
		return
	}
}

func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	var queue [][]int
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}

	queue = append(queue, []int{entrance[0], entrance[1], 0})
	maze[entrance[0]][entrance[1]] = '+'

	for len(queue) > 0 {
		x, y, cur := queue[0][0], queue[0][1], queue[0][2]
		queue = queue[1:]
		if cur > 0 && (x == 0 || x == m-1 || y == 0 || y == n-1) {
			return cur
		}
		for i := 0; i < 4; i++ {
			newX, newY := x+dx[i], y+dy[i]
			if newX >= 0 && newX < m && newY >= 0 && newY < n && maze[newX][newY] == '.' {
				queue = append(queue, []int{newX, newY, cur + 1})
				maze[newX][newY] = '+'
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(nearestExit([][]byte{
		{'+', '+', '.', '+'},
		{'.', '.', '.', '+'},
		{'+', '+', '+', '.'}},
		[]int{0, 2}))
}
