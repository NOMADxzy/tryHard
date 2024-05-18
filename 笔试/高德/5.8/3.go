package main

import "sort"

func carPoolReserve(carpool_travels [][]int, total_seat_count int) bool {
	startCnts := map[int]int{}
	endCnts := map[int]int{}
	keyPoints := map[int]bool{}
	for i := 0; i < len(carpool_travels); i++ {
		e := carpool_travels[i]
		cnt, start, end := e[0], e[1], e[2]
		keyPoints[start] = true
		keyPoints[end] = true
		startCnts[start] += cnt
		endCnts[end] += cnt
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
	return maxLayer <= total_seat_count
}
