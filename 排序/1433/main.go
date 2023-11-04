package main

import (
	"fmt"
	"sort"
)

func checkIfCanBreak(s1 string, s2 string) bool {
	arr1, arr2 := make([]byte, len(s1)), make([]byte, len(s2))
	for i := 0; i < len(s1); i++ {
		arr1[i] = s1[i]
		arr2[i] = s2[i]
	}

	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})

	b1, b2 := false, false
	for i := 0; i < len(s2); i++ {
		if arr1[i] > arr2[i] {
			b1 = true
		} else if arr1[i] < arr2[i] {
			b2 = true
		}
		if b1 && b2 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkIfCanBreak("ace", "bbe"))
}
