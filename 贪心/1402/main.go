package main

import (
	"fmt"
	"sort"
)

func maxSatisfaction(satisfaction []int) int {
	n := len(satisfaction)
	sort.Slice(satisfaction, func(i, j int) bool {
		return satisfaction[i] > satisfaction[j]
	})
	good := 0
	bad := 0
	var i int
	for i = 0; i < n; i++ {
		if satisfaction[i] >= 0 {
			good += satisfaction[i]
		} else {
			if bad-satisfaction[i] >= good {
				break
			}
			bad += -satisfaction[i]
		}
	}
	sum := 0
	for j := i - 1; j >= 0; j-- {
		sum += (i - j) * satisfaction[j]
	}
	return sum
}

func main() {
	fmt.Println(maxSatisfaction([]int{-1, -8, 0, 5, -9}))
}
