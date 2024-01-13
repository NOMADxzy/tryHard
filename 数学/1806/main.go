package main

import "fmt"

func getCommonMul(a, b int) int {
	if a < b {
		a, b = b, a
	}
	mulVal := a * b
	for b > 0 {
		a, b = b, a%b
	}
	return mulVal / a
}

func reinitializePermutation(n int) int {
	common := 1
	for i := 0; i < n; i++ {
		initPos := i
		d := 0
		for d == 0 || initPos != i {
			if i < n/2 {
				i = 2 * i
			} else {
				i = 2*i - n + 1
			}
			d++
		}
		common = getCommonMul(common, d)
	}
	return common
}

func main() {
	fmt.Println(reinitializePermutation(4))
	fmt.Println(getCommonMul(56, 63))
}
