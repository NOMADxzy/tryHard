package main

import "fmt"

func getNext(s string) []int {
	next := make([]int, len(s))
	i, l := 1, 0
	for i < len(s) {
		if s[i] == s[l] {
			l++
			next[i] = l
			i++
		} else {
			if l == 0 {
				next[i] = 0
				i++
			} else {
				l = next[l-1]
			}
		}
	}
	copy(next[1:], next[0:])
	return next
}

func repeatedStringMatch(a string, b string) int {
	next := getNext(b)
	i, j := 0, 0

	for ; j < len(b); i++ {
		if i > len(b)+len(a) {
			return -1
		}
		p := i % len(a)
		if a[p] == b[j] {
			j++
		} else if j == 0 {
		} else {
			j = next[j]
			i--
		}
	}
	return (i-1)/len(a) + 1
}

func main() {
	fmt.Println(repeatedStringMatch("aaaaaaaaaaaaaaaaaaaaaab", "ba"))
}
