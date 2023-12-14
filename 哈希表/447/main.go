package main

import (
	"fmt"
	"math"
)

func numberOfBoomerangs(points [][]int) int {
	m := len(points)
	distMap := map[float64]int{}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			if i != j {
				d := math.Pow(float64(points[i][0]-points[j][0])*float64(points[i][0]-points[j][0])+float64(points[i][1]-points[j][1])*float64(points[i][1]-points[j][1]), 0.5)
				if distMap[d] > 0 {
					ans += 2 * distMap[d]
				}
				distMap[d]++
			}
		}
		clear(distMap)
	}
	return ans
}

func main() {
	fmt.Println(numberOfBoomerangs([][]int{{0, 0}, {1, 0}, {2, 0}}))
}
