package main

import "fmt"

func goNext(arr []int, pos int) bool {
	if arr[pos] == 0 {
		return true
	} else if arr[pos] < 0 {
		return false
	}
	step := arr[pos]
	arr[pos] = -1
	if pos-step >= 0 && goNext(arr, pos-step) {
		return true
	}
	if pos+step < len(arr) && goNext(arr, pos+step) {
		return true
	}
	return false
}

func canReach(arr []int, start int) bool {
	return goNext(arr, start)
}

func main() {
	fmt.Println(canReach([]int{3, 0, 2, 1, 2}, 2))
}
