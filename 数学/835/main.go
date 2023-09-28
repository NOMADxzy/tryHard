package main

import "fmt"

func largestOverlap(img1 [][]int, img2 [][]int) int {
	n := len(img1)
	var OnePoints [][]int

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if img1[i][j] == 1 {
				OnePoints = append(OnePoints, []int{i, j})
			}
		}
	}
	maxVal := 0
	for deltaX := -(n - 1); deltaX <= n-1; deltaX++ {
		for deltaY := -(n - 1); deltaY <= n-1; deltaY++ {
			sum := 0
			for _, point := range OnePoints {
				newX, newY := point[0]+deltaX, point[1]+deltaY
				if newX >= 0 && newX < n && newY >= 0 && newY < n && img2[newX][newY] == 1 {
					sum++
				}
			}
			if sum > maxVal {
				maxVal = sum
			}
		}
	}
	return maxVal
}

func main() {
	fmt.Println()
}
