package main

import "fmt"

func find(n int, acc int) int {
	if n == 1 {
		return acc
	} else if n%2 == 0 {
		return find(n/2, acc+1)
	} else {
		l := find(n-1, acc+1)
		r := find(n+1, acc+1)
		if l < r {
			return l
		} else {
			return r
		}
	}
}

func integerReplacement(n int) int {
	return find(n, 0)
}

func main() {
	fmt.Println(integerReplacement(7))
}
