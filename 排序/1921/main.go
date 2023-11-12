package main

import (
	"fmt"
	"sort"
)

func eliminateMaximum(dist []int, speed []int) int {
	m := len(dist)
	arr := make([]float64, m)

	for i := 0; i < m; i++ {
		arr[i] = float64(dist[i]) / float64(speed[i])
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	ans := 0
	canuseTime := 0
	for i := 0; i < m; i++ {
		if arr[i] > float64(canuseTime) {
			canuseTime++
			ans++
		} else {
			break
		}
	}
	return ans
}

func main() {
	fmt.Println(eliminateMaximum([]int{1, 1, 2, 3}, []int{1, 1, 1, 1}))
}
