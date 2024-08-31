package main

import "fmt"

func solve(s string, t int) int {
	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'M' || s[i] == 'T' {
			cnt++
		}
	}
	return min(len(s), len(s)-cnt+t)
}

func main() {
	m := 0
	t := 0
	var s string
	_, _ = fmt.Scan(&m, &t)
	_, _ = fmt.Scan(&s)
	fmt.Println(solve(s, t))
}
