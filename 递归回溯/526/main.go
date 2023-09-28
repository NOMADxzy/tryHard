package main

import "fmt"

func getCount(i int, n int, state int, history map[int]int) int {
	if history[state] > 0 {
		return history[state]
	}
	if i > n {
		return 1
	}
	mask := 1
	cnt := 0
	for j := 1; j <= n; j++ {
		if state&mask == 0 && (i%j == 0 || j%i == 0) {
			cnt += getCount(i+1, n, state+mask, history)
		}
		mask *= 2
	}
	history[state] = cnt
	return cnt
}

func countArrangement(n int) int {
	history := map[int]int{}
	return getCount(1, n, 0, history)
}

func main() {
	fmt.Println(countArrangement(3))
}
