package main

import "fmt"

func maxScore(cardPoints []int, k int) int {
	sums := make([]int, len(cardPoints))
	sums[0] = cardPoints[0]
	total := cardPoints[0]
	for i := 1; i < len(cardPoints); i++ {
		sums[i] = sums[i-1] + cardPoints[i]
		total += cardPoints[i]
	}
	k = len(cardPoints) - k
	if k == 0 {
		return total
	}
	min := sums[k-1]
	for i := k; i < len(cardPoints); i++ {
		if sums[i]-sums[i-k] < min {
			min = sums[i] - sums[i-k]
		}
	}
	return total - min
}

func main() {

	fmt.Println(maxScore([]int{1, 2, 3, 4, 5, 6, 1}, 3))
}
