package main

import (
	"fmt"
	"sort"
)

func solve(arr []int) int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	m := len(arr)
	dp := make([]int, m)
	last1Idx, last0Idx := -1, -1
	for i := 0; i < m; i++ {
		val := 1
		if arr[i]%2 == 0 {
			if last1Idx >= 0 {
				val += dp[last1Idx]
			}
			last0Idx = i
		} else {
			if last0Idx >= 0 {
				val += dp[last0Idx]
			}
			last1Idx = i
		}
		dp[i] = val
	}
	return dp[m-1]
}

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&arr[i])
	}
	fmt.Println(solve(arr))
}
