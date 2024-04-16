package main

import "fmt"

func solution2(arr []int) int {
	isSu := func(x int) bool {
		for i := 2; i*i <= x; i++ {
			if x%i == 0 {
				return false
			}
		}
		return true
	}
	m := len(arr)
	delCnt := 0
	for i := 0; i < m; i++ {
		if arr[i] == 1 || !isSu(arr[i]) {
			continue
		} else {
			if i+1 < m && isSu(arr[i+1]) {
				delCnt++
				arr[i+1] += arr[i]
			}
		}
	}
	return m - delCnt
}

func main() {
	n := 0
	_, _ = fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&arr[i])
	}
	fmt.Println(solution2(arr))
}
