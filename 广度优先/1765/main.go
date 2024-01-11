package main

import "fmt"

func highestPeak(isWater [][]int) [][]int {
	var queue [][]int
	curH := 1
	m, n := len(isWater), len(isWater[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				isWater[i][j] = 0
				queue = append(queue, []int{i, j})
			} else {
				isWater[i][j] = -1
			}
		}
	}
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for ; len(queue) > 0; curH++ {
		var nextQueue [][]int
		for len(queue) > 0 {
			e := queue[0]
			queue = queue[1:]
			for _, dir := range dirs {
				nx, ny := e[0]+dir[0], e[1]+dir[1]
				if nx >= 0 && nx < m && ny >= 0 && ny < n && isWater[nx][ny] < 0 {
					isWater[nx][ny] = curH
					nextQueue = append(nextQueue, []int{nx, ny})
				}
			}
		}
		queue = nextQueue
	}
	return isWater
}

func main() {
	isWater := [][]int{
		{0, 0, 1},
		{1, 0, 0},
		{0, 0, 0},
	}
	fmt.Println(highestPeak(isWater))
}
