package main

import "sort"

func minCars(intervals [][]int) int {

	startCnts := map[int]int{}
	endCnts := map[int]int{}
	keyPoints := map[int]bool{}
	for i := 0; i < len(intervals); i++ {
		e := intervals[i]
		start, end := e[0], e[1]
		keyPoints[start] = true
		keyPoints[end] = true
		startCnts[start]++
		endCnts[end]++
	}
	maxLayer := 0
	var keyPointsList []int
	for p, _ := range keyPoints {
		keyPointsList = append(keyPointsList, p)
	}
	sort.Slice(keyPointsList, func(i, j int) bool {
		return keyPointsList[i] < keyPointsList[j]
	})
	curLayer := 0
	for _, p := range keyPointsList {
		curLayer += startCnts[p] - endCnts[p]
		if curLayer > maxLayer {
			maxLayer = curLayer
		}
	}
	return maxLayer
}
