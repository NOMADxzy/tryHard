package main

import "fmt"

func carPooling(trips [][]int, capacity int) bool {
	maxLen := 0
	for i := 0; i < len(trips); i++ {
		if trips[i][2] > maxLen {
			maxLen = trips[i][2]
		}
	}
	maxLen++
	deferArr := make([]int, maxLen+1)

	for i := 0; i < len(trips); i++ {
		tp := trips[i]
		deferArr[tp[1]] += tp[0]
		deferArr[tp[2]] -= tp[0]
	}

	if deferArr[0] > capacity {
		return false
	}
	for i := 1; i < len(deferArr); i++ {
		deferArr[i] = deferArr[i-1] + deferArr[i]
		if deferArr[i] > capacity {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(carPooling([][]int{{2, 1, 5}, {3, 5, 7}}, 5))
}
