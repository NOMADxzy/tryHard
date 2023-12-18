package main

import "fmt"

func leastBricks(wall [][]int) int {
	endMap := map[int]int{}
	height := len(wall)
	maxTime := 0
	for _, w := range wall {
		if len(w) == 1 {
			continue
		}
		for i := 0; i < len(w)-1; i++ {
			endMap[w[i]]++
			if endMap[w[i]] > maxTime {
				maxTime = endMap[w[i]]
			}
			w[i+1] += w[i]
		}
	}
	return height - maxTime
}

func main() {
	wall := [][]int{{1, 2, 2, 1}, {3, 1, 2}, {1, 3, 2}, {2, 4}, {3, 1, 2}, {1, 3, 1, 1}}
	fmt.Println(leastBricks(wall))
}
