package main

import "fmt"

func multiply(A int, B int) int {
	ret := 0
	if B == 0 {
		return ret
	}
	if B&1 == 1 {
		ret += A
	}
	ret += 2 * multiply(A, B>>1)
	return ret
}

func main() {
	fmt.Println(multiply(3, 4))
}
