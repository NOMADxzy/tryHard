package main

import "fmt"

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	lowCost := make([]int, n)
	getDist := func(p1, p2 []int) int {
		d1 := p1[0] - p2[0]
		d2 := p1[1] - p2[1]
		return max(d1, -d1) + max(d2, -d2)
	}
	visited := make([]bool, n)
	visited[0] = true
	cost := 0
	for i := 1; i < n; i++ {
		lowCost[i] = getDist(points[0], points[i])
	}
	for i := 1; i < n; i++ {
		best_idx := -1
		best_val := 1 << 31
		for j := 0; j < n; j++ {
			if !visited[j] && lowCost[j] < best_val {
				best_idx = j
				best_val = lowCost[j]
			}
		}
		visited[best_idx] = true
		cost += best_val
		for j := 0; j < n; j++ {
			if !visited[j] {
				if d := getDist(points[best_idx], points[j]); d < lowCost[j] {
					lowCost[j] = d
				}
			}
		}
	}
	return cost
}

func main() {
	points := [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}
	fmt.Println(minCostConnectPoints(points))
}
