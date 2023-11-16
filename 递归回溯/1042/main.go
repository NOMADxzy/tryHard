package main

import "fmt"

func findNext(pos int, neighbors map[int][]int, colors []int) bool {
	if colors[pos-1] > 0 {
		return true
	}
	invalids := make([]bool, 4)
	var waitings []int
	for _, p := range neighbors[pos] {
		if colors[p-1] > 0 {
			invalids[colors[p-1]-1] = true
		} else {
			waitings = append(waitings, p)
		}
	}

	flag := true
	for i := 0; i < 4; i++ {
		if !invalids[i] {
			colors[pos-1] = i + 1
			for _, nextP := range waitings {
				flag = flag && findNext(nextP, neighbors, colors)
				if !flag {
					break
				}
			}
			if flag {
				return true
			}
		}
	}
	return false

}

func gardenNoAdj(n int, paths [][]int) []int {
	neighbors := map[int][]int{}
	for _, path := range paths {
		neighbors[path[0]] = append(neighbors[path[0]], path[1])
		neighbors[path[1]] = append(neighbors[path[1]], path[0])
	}
	colors := make([]int, n)
	for i := 0; i < n; i++ {
		findNext(i+1, neighbors, colors)
	}
	return colors
}

func main() {
	fmt.Println(gardenNoAdj(4, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 3}, {2, 4}}))
}
