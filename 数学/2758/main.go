package main

import "fmt"

func minOperations(n int) int {
	mask := 1
	cnt := 0
	for mask <= n {
		if mask&n > 0 {
			acc := 1
			for (mask*2)&n > 0 {
				mask *= 2
				acc++
			}
			cnt++
			if acc > 1 {
				n += mask * 2
			} else {
				n -= mask
			}
		}
		mask *= 2
	}
	return cnt
}

func main() {
	fmt.Println(minOperations(15))
}
