package main

import (
	"fmt"
	"sort"
)

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	sort.Slice(hand, func(i, j int) bool {
		return hand[i] < hand[j]
	})
	marked := make([]bool, len(hand))
	for i := 0; i < len(hand); i++ {
		if !marked[i] {
			pos := hand[i]
			for j := i; j < len(hand) && pos < hand[i]+groupSize; j++ {
				if !marked[j] && hand[j] == pos {
					pos++
					marked[j] = true
				}
			}
			if pos < hand[i]+groupSize {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3))
}
