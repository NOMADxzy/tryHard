package main

import "fmt"

func partSort(arr []int, start, end int) {
	n := end - start + 1
	if n < 2 {
		arr[start] = 1
		return
	}
	leftCount := (n + 1) / 2
	rightCount := n / 2
	partSort(arr, start, start+leftCount-1)
	partSort(arr, end-rightCount+1, end)
	for i := start; i <= end; i++ {
		if i < start+leftCount {
			arr[i] = arr[i]*2 - 1
		} else {
			arr[i] = arr[i] * 2
		}
	}
}

func beautifulArray(n int) []int {
	arr := make([]int, n)
	partSort(arr, 0, n-1)
	return arr
}

func main() {
	fmt.Println(beautifulArray(5))
}
