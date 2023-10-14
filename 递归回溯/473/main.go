package main

import (
	"fmt"
	"sort"
)

func goNext(sides []int, matchsticks []int, pos int, limit int) bool {
	if pos == len(matchsticks) {
		return true
	}
	for i := 0; i < 4; i++ {
		if sides[i]+matchsticks[pos] <= limit {
			sides[i] += matchsticks[pos]
			if goNext(sides, matchsticks, pos+1, limit) {
				return true
			}
			sides[i] -= matchsticks[pos]
		}
	}
	return false
}

func makesquare(matchsticks []int) bool {
	sides := make([]int, 4)
	limit := 0
	max := 0
	sort.Slice(matchsticks, func(i, j int) bool {
		return matchsticks[i] > matchsticks[j]
	})

	for i := 0; i < len(matchsticks); i++ {
		limit += matchsticks[i]
		if matchsticks[i] > max {
			max = matchsticks[i]
		}
	}
	if limit%4 != 0 || len(matchsticks) < 4 || max > limit/4 {
		return false
	}
	limit /= 4
	return goNext(sides, matchsticks, 0, limit)
}

func main() {
	fmt.Println(makesquare([]int{5969561, 8742425, 2513572, 3352059, 9084275, 2194427, 1017540, 2324577, 6810719, 8936380, 7868365, 2755770, 9954463, 9912280, 4713511}))
}
