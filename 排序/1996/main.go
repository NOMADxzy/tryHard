package main

import (
	"fmt"
	"sort"
)

func numberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		return properties[i][0] < properties[j][0]
	})

	m := len(properties)
	dpRightMax := make([]int, m)
	dpRightMax[m-1] = properties[m-1][1]
	for i := m - 2; i >= 0; i-- {
		dpRightMax[i] = dpRightMax[i+1]
		if properties[i][1] > dpRightMax[i] {
			dpRightMax[i] = properties[i][1]
		}
	}

	ans := 0
	lastPos := m
	for i := m - 2; i >= 0; i-- {
		if properties[i][0] != properties[i+1][0] {
			lastPos = i + 1
		}
		if lastPos < m && dpRightMax[lastPos] > properties[i][1] {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfWeakCharacters([][]int{{5, 5}, {6, 3}, {3, 6}}))
}
