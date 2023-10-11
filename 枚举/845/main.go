package main

import "fmt"

func longestMountain(arr []int) int {
	m := len(arr)
	if m < 3 {
		return 0
	}
	lrArr := make([]int, len(arr))
	rlArr := make([]int, len(arr))

	for i := 1; i < m-1; i++ {
		if arr[i] > arr[i-1] {
			lrArr[i] = lrArr[i-1] + 1
		}
	}
	for i := m - 2; i > 0; i-- {
		if arr[i] > arr[i+1] {
			rlArr[i] = rlArr[i+1] + 1
		}
	}

	ans := 0
	for i := 1; i < m-1; i++ {
		if lrArr[i] > 0 && rlArr[i] > 0 && lrArr[i]+rlArr[i]+1 > ans {
			ans = lrArr[i] + rlArr[i] + 1
		}
	}
	return ans
}

func main() {
	fmt.Println(longestMountain([]int{2, 1, 4, 7, 3, 2, 5}))
}
