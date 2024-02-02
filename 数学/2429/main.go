package main

import "fmt"

func minimizeXor(num1 int, num2 int) int {
	cnt1 := 0
	for i := num2; i > 0; i -= i & (-i) {
		cnt1++
	}
	x := 0
	cnt1Target := 0
	for i := num1; i > 0; i -= i & (-i) {
		cnt1Target++
	}
	mask := 1
	for cnt1 > 0 {
		if num1&mask == 0 {
			if cnt1 > cnt1Target {
				cnt1--
				x += mask
			}
		} else {
			if cnt1 >= cnt1Target {
				cnt1--
				x += mask
			}
			cnt1Target--
		}
		mask *= 2
	}
	return x
}

func main() {
	// 11010    1111
	fmt.Println(minimizeXor(26, 1))
}
