package main

import (
	"fmt"
	"sort"
)

func canReorderDoubled(arr []int) bool {
	m := len(arr)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	nMap := map[int]int{}

	pairCnt := 0
	var i int
	for i = m - 1; i >= 0 && arr[i] >= 0; i-- {
		if cnt, _ := nMap[2*arr[i]]; cnt > 0 {
			nMap[2*arr[i]]--
			pairCnt++
		} else {
			nMap[arr[i]]++
		}
	}
	for ; i >= 0; i-- {
		if arr[i]%2 == 0 && nMap[arr[i]/2] > 0 {
			nMap[arr[i]/2]--
			pairCnt++
		} else {
			nMap[arr[i]]++
		}
	}
	return pairCnt*2 == m
}

func main() {
	fmt.Println(canReorderDoubled([]int{-8, -4, -2, -1, 0, 0, 1, 2, 4, 8}))
}
