package main

import "fmt"

func robotSim(commands []int, obstacles [][]int) int {
	getKey := func(x, y int) int {
		return 60001*(x+30000) + (y + 30000)
	}
	obstaclesMap := map[int]bool{}
	for _, obstacle := range obstacles {
		obstaclesMap[getKey(obstacle[0], obstacle[1])] = true
	}
	x, y := 0, 0
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	changeDir := func(dIdx int, left bool) (idx int) {
		if left {
			idx = dIdx - 1
			if idx < 0 {
				idx = 3
			}
		} else {
			idx = dIdx + 1
			if idx > 3 {
				idx = 0
			}
		}
		return
	}
	bestDist := 0
	dIdx := 0
	for _, c := range commands {
		if c < 0 {
			dIdx = changeDir(dIdx, c == -2)
		} else {
			for c > 0 {
				nx := x + dirs[dIdx][0]
				ny := y + dirs[dIdx][1]
				if obstaclesMap[getKey(nx, ny)] {
					break
				}
				x, y = nx, ny
				c--
			}
			bestDist = max(x*x+y*y, bestDist)
		}
	}
	return bestDist
}

func main() {
	fmt.Println(robotSim([]int{4, -1, 4, -2, 4}, [][]int{{2, 4}}))
}
