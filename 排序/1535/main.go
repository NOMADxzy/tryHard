package main

import "fmt"

func getWinner(arr []int, k int) int {
	cnt := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
			cnt++
		} else {
			cnt = 1
		}
		if cnt == k {
			return arr[i+1]
		}
	}
	return arr[len(arr)-1]
}

func main() {
	fmt.Println(getWinner([]int{2, 1, 3, 5, 4, 6, 7}, 1))
}
