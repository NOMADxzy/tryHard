package main

import (
	"fmt"
	"math"
)

func partQuickSort(arr [][]float64, start, end int, k int) {
	if end-start < 1 {
		return
	}
	guard := arr[start][0]
	l, r := start, end
	for l < r {
		for l < r && arr[r][0] >= guard {
			r--
		}
		arr[r], arr[l] = arr[l], arr[r]
		for l < r && arr[l][0] <= guard {
			l++
		}
		arr[r], arr[l] = arr[l], arr[r]
	}
	if l == k {
		return
	} else if l < k {
		partQuickSort(arr, l+1, end, k)
	} else {
		partQuickSort(arr, start, l-1, k)
	}
}

func kClosest(points [][]int, k int) [][]int {
	dists := make([][]float64, len(points))
	for i := 0; i < len(points); i++ {
		dist := math.Pow(float64(points[i][0]*points[i][0]+points[i][1]*points[i][1]), 0.5)
		dists[i] = []float64{dist, float64(i)}
	}
	partQuickSort(dists, 0, len(dists)-1, k)
	ans := make([][]int, k)
	for i := 0; i < k; i++ {
		ans[i] = points[int(dists[i][1])]
	}
	return ans
}

func main() {
	points := [][]int{{3, 3}, {5, -1}, {-2, 4}}
	fmt.Println(kClosest(points, 2))
}
