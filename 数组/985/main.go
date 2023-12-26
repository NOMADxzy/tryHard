package main

import "fmt"

func sumEvenAfterQueries(nums []int, queries [][]int) []int {
	m := len(nums)
	sum := 0
	for i := 0; i < m; i++ {
		if nums[i]%2 == 0 {
			sum += nums[i]
		}
	}
	ans := make([]int, len(queries))
	for i := 0; i < len(ans); i++ {
		v, idx := queries[i][0], queries[i][1]
		pre := nums[idx]
		nums[idx] += v
		if pre%2 == 0 {
			sum -= pre
		}
		if nums[idx]%2 == 0 {
			sum += nums[idx]
		}
		ans[i] = sum
	}
	return ans
}

func main() {
	fmt.Println(sumEvenAfterQueries([]int{1, 2, 3, 4}, [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}}))
}
