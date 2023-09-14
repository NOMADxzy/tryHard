package main

import "fmt"

func minFlips(a int, b int, c int) int {
	mask := 1
	cnt := 0
	for mask <= a || mask <= b || mask <= c {
		x1 := a & mask
		x2 := b & mask
		if c&mask == 0 {
			cnt += x1/mask + x2/mask
		} else {
			if x1 == 0 && x2 == 0 {
				cnt++
			}
		}
		mask *= 2
	}
	return cnt
}
func main() {
	fmt.Println(minFlips(8, 3, 5))
}
