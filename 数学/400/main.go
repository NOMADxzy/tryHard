package main

import "fmt"

func findNthDigit(n int) int {

	var width int
	pow := 1
	for width = 1; n > 0; width++ {
		n -= width * 9 * pow
		pow *= 10
	}
	width--
	pow /= 10
	n += width * 9 * pow

	base := pow
	k := n / width
	j := n % width
	if j == 0 {
		k--
		j = width
	}
	base += k

	c := width - j + 1

	res := 0
	for c > 0 {
		res = base % 10
		base /= 10
		c--
	}
	return res
}

func main() {
	fmt.Println(findNthDigit(100))
}
