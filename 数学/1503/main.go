package main

import "fmt"

func getLastMoment(n int, left []int, right []int) int {
	leftR, rightL := 0, n
	for i := 0; i < len(left); i++ {
		if left[i] > leftR {
			leftR = left[i]
		}
	}
	for i := 0; i < len(right); i++ {
		if right[i] < rightL {
			rightL = right[i]
		}
	}
	return max(n-rightL, leftR)
}
func main() {
	fmt.Println(getLastMoment(4, []int{4, 3}, []int{0, 1}))
}
