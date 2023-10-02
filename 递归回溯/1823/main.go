package main

import "fmt"

func findNext(mark []bool, start, k int, n int, remain int) int {
	if remain == 1 {
		return start
	}
	var cnt, i int
	for i = start - 1; cnt < k; {
		i++
		if i > n {
			i = 1
		}
		if !mark[i-1] {
			cnt++
		}
	}
	mark[i-1] = true
	for mark[i-1] {
		i++
		if i > n {
			i = 1
		}
	}
	return findNext(mark, i, k, n, remain-1)
}

func findTheWinner(n int, k int) int {
	mark := make([]bool, n)
	return findNext(mark, 1, k, n, n)
}

func main() {
	fmt.Println(findTheWinner(6, 5))
}
