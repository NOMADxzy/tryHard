package main

import "fmt"

func rangeBitwiseAnd(left int, right int) int {

	sum := 0
	tailMask := 0

	for mask := 1; mask <= left; mask *= 2 {
		cur := left & mask
		if cur > 0 {
			k := left + mask - left&tailMask
			if k > right {
				sum = left - left&tailMask
				break
			}
		}
		tailMask += mask
	}
	return sum
}

func main() {
	fmt.Println(rangeBitwiseAnd(5, 8))
}
