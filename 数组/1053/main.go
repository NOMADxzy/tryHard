package main

import "fmt"

func prevPermOpt1(arr []int) []int {
	minRightIdx := len(arr) - 1
	var i int
	for i = len(arr) - 2; i >= 0; i-- {
		if arr[i] <= arr[minRightIdx] {
			minRightIdx = i
		} else {
			break
		}
	}
	if i < 0 {
		return arr
	}
	left := i
	right := minRightIdx
	for i < len(arr) {
		if arr[i] < arr[left] && arr[i] > arr[right] {
			right = i
		}
		i++
	}
	arr[left], arr[right] = arr[right], arr[left]
	return arr
}

func main() {
	fmt.Println(prevPermOpt1([]int{1, 1, 5}))
}
