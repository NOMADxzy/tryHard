package main

import (
	"fmt"
	"sort"
)

func findW(n int, history map[int]int) int {
	if v, ok := history[n]; ok {
		return v
	}
	var res int
	if n%2 == 0 {
		res = findW(n/2, history) + 1
	} else {
		res = findW(n*3+1, history) + 1
	}
	history[n] = res
	return res
}

func getKth(lo int, hi int, k int) int {
	m := hi - lo + 1
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, 2)
		dp[i][1] = lo + i
	}

	history := map[int]int{1: 0}
	for i := lo; i <= hi; i++ {
		dp[i-lo][0] = findW(i, history)
	}
	sort.Slice(dp, func(i, j int) bool {
		return dp[i][0] < dp[j][0]
	})

	w := dp[k-1][0]
	acc := 0
	var arr []int
	for i := 0; i < len(dp); i++ {
		if dp[i][0] < w {
			acc++
		} else if dp[i][0] == w {
			arr = append(arr, dp[i][1])
		}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	return arr[k-1-acc]
}

func main() {
	fmt.Println(getKth(1, 1000, 777))
}
