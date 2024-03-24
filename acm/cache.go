package main

import (
	"fmt"
)

func solve(arr []int) string {
	sum := 1
	cnt0, cnt1 := 0, 0
	tol := 0
	for i := 0; i < len(arr); i++ {
		sum *= arr[i]
		if arr[i]%2 == 0 {
			cnt0++
		} else {
			cnt1++
		}
		for j := arr[i]; j > 1; j /= 2 {
			if j%2 == 1 {
				tol++
			}
		}
	}
	if tol%len(arr) == 0 {
		return "YES"
	} else {
		return "NO"
	}

}

func main() {
	var T, tmp, n int
	_, _ = fmt.Scan(&T)
	for T > 0 {
		_, _ = fmt.Scan(&n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			_, _ = fmt.Scan(&tmp)
			arr[i] = tmp
		}
		fmt.Println(solve(arr))
		T--
	}
}
