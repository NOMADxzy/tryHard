package main

import (
	"fmt"
)

func solve(arr []int, d int) int {
	m := len(arr)
	preMinIdx, nextMaxIdx := make([]int, m), make([]int, m)
	var l, r int
	for r < m {
		for r+1 < m && arr[l]+d >= arr[r+1] {
			r++
		}
		nextMaxIdx[l] = r
		l++
		if l > r {
			r++
		}
	}
	for l < m {
		nextMaxIdx[l] = m - 1
	}
	l, r = m-1, m-1
	for l >= 0 {
		for l-1 >= 0 && arr[r]-d <= arr[l-1] {
			l--
		}
		preMinIdx[r] = l
		r--
		if l > r {
			l--
		}
	}
	for r >= 0 {
		preMinIdx[r] = 0
	}
	// fmt.Println(preMinIdx)
	// fmt.Println(nextMaxIdx)
	ans := 0
	for i := 0; i < m; i++ {
		if v := nextMaxIdx[i] - i; v >= 2 {
			ans = (ans + v*(v-1)/2) % 99997867
		}
		// ans += (i - preMinIdx[i]) * (nextMaxIdx[i] - i)
	}
	return ans
}

func main() {
	N := 0
	D := 0
	_, _ = fmt.Scan(&N, &D)
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		_, _ = fmt.Scan(&arr[i])
	}
	fmt.Println(solve(arr, D))
}
