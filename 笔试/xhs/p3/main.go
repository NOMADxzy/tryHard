package main

import "fmt"

func solve(arr []int) []int {
	var sumVal, maxVal int
	maxVal = arr[0]
	m := len(arr)
	for i := 0; i < m; i++ {
		sumVal += arr[i]
		if maxVal < arr[i] {
			maxVal = arr[i]
		}
	}
	ans := make([]int, m)
	for i := 0; i < m; i++ {
		if m == 2 && arr[i]-arr[1-i] > 1 {
			ans[i] = -1
			continue
		}
		target := maxVal
		for target-arr[i] > (m-1)*target-(sumVal-arr[i])+1 {
			target++
		}
		ans[i] = target + (sumVal - arr[i])
		if target-arr[i]-1 > 0 {
			ans[i] += target - arr[i] - 1
		}
	}
	return ans
}

func main() {
	var n, tmp int
	_, _ = fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < len(arr); i++ {
		_, _ = fmt.Scan(&tmp)
		arr[i] = tmp
	}
	ans := solve(arr)
	for _, an := range ans {
		fmt.Println(an)
	}
}
