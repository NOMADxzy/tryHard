package main

import "fmt"

func solve(n int) [][]int {
	var ans [][]int
	for i := 0; i < n; i++ {
		id1 := i + 1
		for j := i + 1; j < n; j++ {
			id2 := j + 1
			ans = append(ans, []int{id1, id2})
		}
	}
	return ans
}

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	ans := solve(n)
	for _, an := range ans {
		fmt.Println(an[0], an[1])
	}
}
