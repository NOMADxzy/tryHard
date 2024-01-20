package main

import "fmt"

func winnerOfGame(colors string) bool {
	acnts, bcnts := 0, 0
	for i := 0; i < len(colors); {
		if colors[i] == 'A' {
			pre := i
			for i < len(colors) && colors[i] == 'A' {
				i++
			}
			acnts += max(0, i-pre-2)
		} else {
			pre := i
			for i < len(colors) && colors[i] == 'B' {
				i++
			}
			bcnts += max(0, i-pre-2)
		}
	}
	return acnts > bcnts
}

func main() {
	fmt.Println(winnerOfGame("ABBBBBBBAAA"))
}
