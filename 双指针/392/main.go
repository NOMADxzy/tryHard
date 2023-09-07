package main

import "fmt"

func isSubsequence(s string, t string) bool {
	p := 0
	q := 0
	m, n := len(s), len(t)

	for p < m && q < n {
		if s[p] == t[q] {
			p++
			q++
		} else {
			q++
		}
	}
	return p == m
}

func main() {
	fmt.Println(isSubsequence("aec", "abcde"))
}
