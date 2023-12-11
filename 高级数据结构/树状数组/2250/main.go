package main

import (
	"fmt"
	"sort"
)

func lowBit(x int) int {
	return x & (-x)
}

func update(tree []int, x int, val int) {
	for x < len(tree) {
		tree[x] += val
		x += lowBit(x)
	}
}

func query(tree []int, x int) {
	sum := 0
	for x > 0 {
		sum += tree[x]
		x -= lowBit(x)
	}
}

// 朴素解法
func countRectangles1(rectangles [][]int, points [][]int) []int {
	cntRect := len(rectangles)
	maxY, maxX := 0, 0

	for i := 0; i < len(rectangles); i++ {
		if rectangles[i][0] > maxX {
			maxX = rectangles[i][0]
		}
		if rectangles[i][1] > maxY {
			maxY = rectangles[i][1]
		}
	}
	maxX++
	maxY++
	history := make([][]int, maxX)
	for i := 0; i < maxX; i++ {
		history[i] = make([]int, maxY)
	}
	for i := 0; i < cntRect; i++ {
		rect := rectangles[i]
		for j := 1; j <= rect[0]; j++ {
			for k := 1; k <= rect[1]; k++ {
				history[j][k]++
			}
		}
	}
	res := make([]int, len(points))
	for i := 0; i < len(res); i++ {
		x, y := points[i][0], points[i][1]
		if x >= maxX || y >= maxY {
			res[i] = 0
			continue
		}
		res[i] = history[x][y]
	}
	return res
}

func countRectangles(rectangles [][]int, points [][]int) []int {
	rectsPerHeight := make([][]int, 101)
	for i := 0; i < len(rectangles); i++ {
		rect := rectangles[i]
		rectsPerHeight[rect[1]] = append(rectsPerHeight[rect[1]], rect[0])
	}
	for i := 1; i < 101; i++ {
		if len(rectsPerHeight[i]) > 0 {
			sort.Slice(rectsPerHeight[i], func(j, k int) bool {
				return rectsPerHeight[i][j] < rectsPerHeight[i][k]
			})
		}
	}
	res := make([]int, len(points))
	for i := 0; i < len(points); i++ {
		x, y := points[i][0], points[i][1]
		cnt := 0
		for j := y; j < 101; j++ {
			arr := rectsPerHeight[j]
			if len(arr) > 0 {
				l, r := 0, len(arr)-1
				if arr[0] >= x {
					cnt += len(arr)
					continue
				} else if arr[r] < x {
					continue
				} else {
					for l < r {
						mid := (l + r) / 2
						if arr[mid] < x {
							l = mid + 1
						} else {
							r = mid
						}
					}
					cnt += len(arr) - r
				}
			}
		}
		res[i] = cnt
	}
	return res
}
func main() {
	fmt.Println(countRectangles([][]int{{1, 2}, {2, 3}, {2, 5}}, [][]int{{2, 1}, {1, 4}}))
}
