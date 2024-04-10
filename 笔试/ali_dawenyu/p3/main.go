package main

import (
	"fmt"
)

func solve(s string, m int) int {
	sum := 0
	prePos := int(s[0] - 'A')
	getDist := func(a, b int) int {
		if a > b {
			a, b = b, a
		}
		v := b - a
		if a+26-b < v {
			v = a + 26 - b
		}
		return v
	}
	n := len(s)
	distArr := make([]int, n)
	for i := 0; i < len(s); i++ {
		targetPos := int(s[i] - 'A')
		distArr[i] = getDist(prePos, targetPos) + 1
		sum += distArr[i]
		prePos = targetPos
	}
	if m > n {
		return sum
	}
	windowVal := 0
	maxWindowVal := 0
	for i := 0; i < m; i++ {
		windowVal += distArr[i]
	}
	for i := m; i < n; i++ {
		windowVal += distArr[i]
		windowVal -= distArr[i-m]
		if windowVal > maxWindowVal {
			maxWindowVal = windowVal
		}
	}
	newSum := sum - maxWindowVal + m*2
	if newSum < sum {
		sum = newSum
	}
	return sum
}

func main() {
	var s string
	var m int
	_, _ = fmt.Scan(&s)
	_, _ = fmt.Scan(&m)
	fmt.Println(solve(s, m))
}
