package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	y1 int
	y2 int
}

func minAreaRect(points [][]int) int {
	m := len(points)
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0] || points[i][0] == points[j][0] && points[i][1] < points[j][1]
	})
	best := 1 << 31

	pairMap := map[Pair]int{}

	pos := 0
	for pos < m {
		i, j := pos, pos
		for j+1 < m && points[j+1][0] == points[i][0] {
			j++
		}
		for l := i; l < j; l++ {
			for r := l + 1; r <= j; r++ {
				if h, ok := pairMap[Pair{points[l][1], points[r][1]}]; ok {
					s := (points[r][1] - points[l][1]) * (points[pos][0] - h)
					if s < best {
						best = s
					}
				}
				pairMap[Pair{points[l][1], points[r][1]}] = points[pos][0]
			}
		}
		pos = j + 1
	}
	if best == 1<<31 {
		return 0
	}
	return best

}

func main() {
	fmt.Println(minAreaRect([][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {2, 2}}))
}
