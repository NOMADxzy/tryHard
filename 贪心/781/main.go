package main

import (
	"fmt"
	"sort"
)

func numRabbits(answers []int) int {
	ans := 0
	tmp := 0
	last := 0
	sort.Slice(answers, func(i, j int) bool {
		return answers[i] < answers[j]
	})

	for i := 0; i < len(answers); i++ {
		if answers[i] == last && tmp > 0 {
			tmp--
		} else {
			ans += answers[i] + 1
			tmp = answers[i]
			last = answers[i]
		}
	}
	return ans
}

func main() {
	fmt.Println(numRabbits([]int{0, 10, 10, 0}))
}
