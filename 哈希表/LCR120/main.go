package main

import "fmt"

func findRepeatDocument(documents []int) int {
	m := len(documents)
	for i := 0; i < m; i++ {
		if i != documents[i] {
			targetPlace := documents[i]
			if documents[targetPlace] == targetPlace {
				return targetPlace
			} else {
				documents[targetPlace], documents[i] = documents[i], documents[targetPlace]
				i--
			}
		}
	}
	return m
}

func main() {
	fmt.Println(findRepeatDocument([]int{2, 5, 3, 0, 5, 0}))
}
