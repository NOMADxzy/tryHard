package main

import (
	"fmt"
	"sort"
)

func bagOfTokensScore(tokens []int, power int) int {
	sort.Slice(tokens, func(i, j int) bool {
		return tokens[i] < tokens[j]
	})
	n := len(tokens)
	left, right := 0, n-1
	score := 0
	for left <= right {
		if left == right {
			if power >= tokens[left] {
				score++
			}
			break
		}
		if power > tokens[left] {
			for power >= tokens[left] {
				power -= tokens[left]
				score++
				left++
			}
		} else if score == 0 {
			break
		}
		if left < right && score > 0 {
			score--
			power += tokens[right]
			right--
		}
	}
	return score
}

func main() {
	fmt.Println(bagOfTokensScore([]int{100, 200, 300, 400}, 200))
}
