package main

import "fmt"

func getDescentPeriods(prices []int) int64 {
	m := len(prices)
	ans := int64(1)
	pre := int64(1)
	var cur int64

	for i := 1; i < m; i++ {
		cur = 1
		if prices[i] == prices[i-1]-1 {
			cur += pre
		}
		pre = cur
		ans += cur
	}
	return ans
}

func main() {
	fmt.Println(getDescentPeriods([]int{3, 2, 1, 4}))
}
