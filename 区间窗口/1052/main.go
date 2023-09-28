package main

import "fmt"

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	score := 0
	l, r := 1, minutes
	for i := 0; i < len(customers); i++ {
		if i < r || grumpy[i] == 0 {
			score += customers[i]
		}
	}
	maxScore := score
	for r < len(customers) {
		if grumpy[r] == 1 {
			score += customers[r]
		}
		if grumpy[l-1] == 1 {
			score -= customers[l-1]
		}
		if score > maxScore {
			maxScore = score
		}
		l++
		r++
	}
	return maxScore
}

func main() {
	customers := []int{1, 0, 1, 2, 1, 1, 7, 5}
	grumpy := []int{0, 1, 0, 1, 0, 1, 0, 1}
	fmt.Println(maxSatisfied(customers, grumpy, 2))
}
