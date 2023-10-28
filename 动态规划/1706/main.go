package main

import "fmt"

func findBall(grid [][]int) []int {
	m, n := len(grid), len(grid[0])

	balls := make([]int, n)
	for i := 0; i < n; i++ {
		balls[i] = i
	}

	for r := 0; r < m; r++ {
		for idx, ballPos := range balls {
			if ballPos < 0 {
				continue
			}
			if grid[r][ballPos] == 1 && ballPos < n-1 && grid[r][ballPos+1] == 1 {
				balls[idx]++
			} else if grid[r][ballPos] == -1 && ballPos > 0 && grid[r][ballPos-1] == -1 {
				balls[idx]--
			} else {
				balls[idx] = -1
			}
		}
		fmt.Println(balls)
	}

	return balls
}

func main() {
	grid := [][]int{{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1}}
	fmt.Println(findBall(grid))
}
