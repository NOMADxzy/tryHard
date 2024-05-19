package main

import (
	"fmt"
	"sort"
)

func rectangleArea(rectangles [][]int) int {
	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i][0] < rectangles[j][0]
	})
	MOD := 1000000007
	keyPoints := map[int]bool{}
	var keyPointsList []int
	var segment_start []int
	var segment_end []int
	var segment_cnt []int
	var sweep [][]int
	for i, rectangle := range rectangles {
		if keyPoints[rectangle[1]] == false {
			keyPoints[rectangle[1]] = true
			keyPointsList = append(keyPointsList, rectangle[1])
		}
		if keyPoints[rectangle[3]] == false {
			keyPoints[rectangle[3]] = true
			keyPointsList = append(keyPointsList, rectangle[3])
		}
		sweep = append(sweep, []int{rectangle[0], i, 1})
		sweep = append(sweep, []int{rectangle[2], i, -1})
	}
	sort.Slice(sweep, func(i, j int) bool {
		return sweep[i][0] < sweep[j][0]
	})
	sort.Slice(keyPointsList, func(i, j int) bool {
		return keyPointsList[i] < keyPointsList[j]
	})
	for i := 0; i < len(keyPointsList)-1; i++ {
		segment_start = append(segment_start, keyPointsList[i])
		segment_end = append(segment_end, keyPointsList[i+1])
		segment_cnt = append(segment_cnt, 0)
	}
	sum := 0
	ans := 0

	for idx, pair := range sweep {
		rectangle := rectangles[pair[1]]
		mul := pair[2]
		_, y1, _, y2 := rectangle[0], rectangle[1], rectangle[2], rectangle[3]
		for i := 0; i < len(segment_start); i++ {
			if segment_start[i] >= y1 {
				if segment_end[i] <= y2 {
					segment_cnt[i] += mul
					if segment_cnt[i] == 1 && mul == 1 {
						sum += segment_end[i] - segment_start[i]
					} else if segment_cnt[i] == 0 && mul == -1 {
						sum -= segment_end[i] - segment_start[i]
					}
				} else {
					break
				}
			}
		}
		if idx < len(sweep)-1 {
			ans = (ans + (sum*(sweep[idx+1][0]-sweep[idx][0]))%MOD) % MOD
		}
	}
	return ans
}

func main() {
	rects := [][]int{
		{0, 0, 2, 2},
		{1, 0, 2, 3},
		{1, 0, 3, 1},
	}
	fmt.Println(rectangleArea(rects))
}
