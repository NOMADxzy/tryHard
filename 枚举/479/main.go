package main

import "fmt"

func get10Mul(x int) int {
	r := 1
	for ; x > 0; x-- {
		r *= 10
	}
	return r
}

func largestPalindrome(n int) int {
	for w := 2 * n; w >= n; w-- {
		leftW := w / 2
		maxLeft, minLeft := 0, 0
		if leftW > 0 {
			maxLeft, minLeft = 9, 1
			for i := 0; i < leftW-1; i++ {
				maxLeft = 10*maxLeft + 9
				minLeft = 10 * minLeft
			}
		}
		for left := maxLeft; left >= minLeft; left-- {
			right := 0
			for i := left; i > 0; i /= 10 {
				right = 10*right + i%10
			}
			mid := w - 2*leftW
			if mid > 0 {
				for j := 9; j >= 0; j-- {
					target := right + j*get10Mul(leftW) + left*get10Mul(w-leftW)
					for x := get10Mul(n) - 1; x*x >= target; x-- {
						if target%x == 0 {
							return target % 1337
						}
					}
				}
			} else {
				target := right + left*get10Mul(w-leftW)
				for x := get10Mul(n) - 1; x*x >= target; x-- {
					if target%x == 0 {
						return target % 1337
					}
				}
			}

		}
	}
	return 0
}

func main() {
	fmt.Println(largestPalindrome(1))
}
