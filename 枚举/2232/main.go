package main

import (
	"fmt"
	"strconv"
)

func minimizeResult(expression string) string {
	var pos int
	for i := 0; i < len(expression); i++ {
		if expression[i] == '+' {
			pos = i
			break
		}
	}
	leftSize, rightSize := pos, len(expression)-1-pos
	leftCombines := make([][]int, leftSize)
	rightCombines := make([][]int, rightSize)
	leftArr := make([]int, leftSize)
	rightArr := make([]int, rightSize)

	for i := 0; i < leftSize; i++ {
		leftArr[i] = int(expression[i] - '0')
	}
	for i := pos + 1; i < len(expression); i++ {
		rightArr[i-pos-1] = int(expression[i] - '0')
	}

	L := 0
	R, _ := strconv.Atoi(expression[:leftSize])
	pow := 1
	for i := 1; i < leftSize; i++ {
		pow *= 10
	}
	for i := 0; i < pos; i++ {
		leftCombines[i] = []int{L, R}
		L = L*10 + leftArr[i]
		R -= leftArr[i] * pow
		pow /= 10
	}

	L = 0
	R, _ = strconv.Atoi(expression[pos+1:])
	pow = 1
	for i := 1; i < rightSize; i++ {
		pow *= 10
	}
	for i := 0; i < rightSize; i++ {
		L = L*10 + rightArr[i]
		R -= rightArr[i] * pow
		rightCombines[i] = []int{L, R}
		pow /= 10
	}

	ans := 10000000000
	lb, rb := 0, 0
	for i := 0; i < len(leftCombines); i++ {
		for j := 0; j < len(rightCombines); j++ {
			left := leftCombines[i][0]
			right := rightCombines[j][1]
			mid := leftCombines[i][1] + rightCombines[j][0]
			if left == 0 {
				left = 1
			}
			if right == 0 {
				right = 1
			}
			if left*mid*right < ans {
				ans = left * mid * right
				lb, rb = i, j
			}
		}
	}
	s1, s2, s3 := "", expression[lb:pos+rb+2], expression[pos+rb+2:]
	if lb > 0 {
		s1 = expression[:lb]
	}
	return s1 + "(" + s2 + ")" + s3
}

func main() {
	fmt.Println(minimizeResult("247+38"))
}
