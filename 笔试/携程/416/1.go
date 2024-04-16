package main

import (
	"fmt"
)

func solution1(s string) int {
	cnt := 0
	i := 0
	for i+1 < len(s) {
		if s[i] == 'y' && s[i+1] == 'u' {
			cnt++
			i += 2
		} else {
			i++
		}
	}
	return cnt
}

func main() {
	var s string
	_, _ = fmt.Scan(&s)
	fmt.Println(solution1(s))
}
