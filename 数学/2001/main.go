package main

import "fmt"

func interchangeableRectangles(rectangles [][]int) int64 {
	var ans int64
	getCommonMul := func(a, b int) int {
		if a < b {
			a, b = b, a
		}
		for b > 0 {
			a, b = b, a%b
		}
		return a
	}
	cnts := map[int]int{}
	for _, rectangle := range rectangles {
		x, y := rectangle[0], rectangle[1]
		m := getCommonMul(x, y)
		x /= m
		y /= m
		ans += int64(cnts[10001*x+y])
		cnts[10001*x+y]++
	}
	return ans
}

func main() {
	fmt.Println(interchangeableRectangles([][]int{{4, 8}, {3, 6}, {10, 20}, {15, 30}}))
}
