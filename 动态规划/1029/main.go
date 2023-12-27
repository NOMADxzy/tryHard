package main

import "fmt"

func twoCitySchedCost(costs [][]int) int {
	n := len(costs) / 2
	preCosts := make([]int, 2)
	preCosts[0], preCosts[1] = costs[0][1], costs[0][0]
	lastMinA := 0
	for i := 1; i < len(costs); i++ {
		cur := costs[i]
		minA, maxA := max(0, i+1-n), min(n, i+1)
		newCosts := make([]int, maxA-minA+1)
		for j := 0; j < len(newCosts); j++ {
			if j+minA == 0 {
				newCosts[j] = preCosts[0] + cur[1]
			} else if j+minA-lastMinA < len(preCosts) {
				newCosts[j] = min(preCosts[j-1+minA-lastMinA]+cur[0], preCosts[j+minA-lastMinA]+cur[1])
			} else {
				newCosts[j] = preCosts[j-1+minA-lastMinA] + cur[0]
			}
		}
		preCosts = newCosts
		lastMinA = minA
	}
	return preCosts[len(preCosts)-1]
}

func main() {
	fmt.Println(twoCitySchedCost([][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}}))
}
