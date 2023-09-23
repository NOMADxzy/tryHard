package main

import "fmt"

func lastRemaining(n int) int {
	step := 2
	dir := 1
	//mark := make([]bool, n)
	cnt := 0
	pos := 0
	for cnt < n-1 {
		for ; pos < n && pos >= 0; pos += step * dir {
			cnt++
		}
		pos -= step * dir
		if pos+dir*step/2 >= 0 && pos+dir*step/2 < n {
			pos += dir * step / 2
		} else {
			pos -= dir * step / 2
		}
		dir = -dir
		step *= 2
	}
	return pos + 1
}

func main() {
	fmt.Println(lastRemaining(1))
}
