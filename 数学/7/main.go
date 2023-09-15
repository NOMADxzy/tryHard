package main

import (
	"fmt"
	"strconv"
)

func reverse(x int) int {
	max_int := 2147483647
	neg := false
	if x < 0 {
		neg = true
		max_int++
		x = -x
	}
	s := strconv.Itoa(x)
	m := len(s)
	power := 1
	acc := 0
	for i := 0; i < m; i++ {
		k := int(s[i] - '0')
		if max_int-acc >= k*power {
			acc += k * power
		} else {
			return 0
		}
		power *= 10
	}
	if neg {
		acc = -acc
	}
	return acc
}

func main() {
	fmt.Println(reverse(1534236461))
}
