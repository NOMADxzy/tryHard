package main

import "fmt"

func cntNext(leftChild, rightChild []int, pos int, mark []bool, hist map[int]int) int {
	if pos >= 0 && mark[pos] {
		return -1 << 31
	} else if val, ok := hist[pos]; ok {
		return val
	}
	mark[pos] = true
	hist[pos] = cntNext(leftChild, rightChild, leftChild[pos], mark, hist) + cntNext(leftChild, rightChild, rightChild[pos], mark, hist) + 1
	return hist[pos]
}

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	hist := map[int]int{-1: 0}
	for i := 0; i < n; i++ {
		if cntNext(leftChild, rightChild, i, make([]bool, n), hist) == n {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(validateBinaryTreeNodes(4, []int{1, -1, 3, -1}, []int{2, -1, -1, -1}))
}
