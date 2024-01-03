package main

import "fmt"

func circularPermutation(n int, start int) []int {
	masks := make([]int, 30)
	mask := 1
	for i := 0; i < len(masks); i++ {
		masks[i] = mask
		mask *= 2
	}
	var partition func(n int) []int
	partition = func(n int) []int {
		if n == 1 {
			return []int{0, 1}
		}
		left := partition(n - 1)
		newArr := make([]int, 2*len(left))
		for i := 0; i < len(left); i++ {
			newArr[i] = left[i]
			newArr[len(newArr)-1-i] = masks[n-1] + left[i]
		}
		return newArr
	}
	ans := partition(n)
	fans := make([]int, len(ans))
	var mid int
	for ans[mid] != start {
		mid++
	}
	copy(fans, ans[mid:])
	copy(fans[len(fans)-mid:], ans[:mid])
	return fans
}

func main() {
	fmt.Println(circularPermutation(3, 0))
}
