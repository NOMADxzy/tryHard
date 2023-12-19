package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	root := math.Pow(float64(c), 0.5)
	right := int(root) + 1
	left := 0

	for left <= right {
		v := left*left + right*right
		if v < c {
			left++
		} else if v > c {
			right--
		} else {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(judgeSquareSum(26))
}
