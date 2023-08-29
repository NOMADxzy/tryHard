package main

import "fmt"

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func hIndex(citations []int) int {
	H := 0
	var cits []int

	for i := 0; i < len(citations); i++ {
		for j := 0; j < citations[i]; j++ {
			if j < len(cits) {
				cits[j] += 1
			} else {
				for j >= len(cits) {
					cits = append(cits, 1)
				}
			}

		}
	}
	for i, c := range cits {
		k := min(i+1, c)
		if k > H {
			H = k
		}
	}
	return H
}

func main() {
	fmt.Println(hIndex([]int{0, 1}))
}
