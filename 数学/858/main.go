package main

import "fmt"

func mirrorReflection(p int, q int) int {
	var a []int
	for p%2 == 0 && q%2 == 0 {
		p /= 2
		q /= 2
	}
	if p%2 == 0 {
		a = []int{-1, 2}
	} else {
		a = []int{0, 1}
	}
	return a[q%2]
}

func main() {
	fmt.Println(mirrorReflection(4, 2))
}
