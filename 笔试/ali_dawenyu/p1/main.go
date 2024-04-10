package main

import (
	"fmt"
)

func solve(s string) string {
	lowBit := func(x int) int {
		return x & (-x)
	}
	ans := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		cnt := 0
		for j := i + 1; j > 0; j -= lowBit(j) {
			cnt++
		}
		ans[i] = s[i]
		if cnt%2 == 1 {
			ans[i] = byte(ans[i] - 32)
		}
	}
	return string(ans)
}

func main() {
	var T int
	var s string
	_, _ = fmt.Scan(&T)

	for i := 0; i < T; i++ {
		_, _ = fmt.Scan(&s)
		fmt.Println(solve(s))
	}
}
