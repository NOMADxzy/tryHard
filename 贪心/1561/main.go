package main

import (
	"fmt"
	"sort"
)

func maxCoins(piles []int) int {
	ans := 0
	n := len(piles) / 3
	sort.Slice(piles, func(i, j int) bool {
		return piles[i] > piles[j]
	})
	for i := 0; i < n; i++ {
		ans += piles[2*i+1]
	}
	return ans
}

func main() {
	fmt.Println(maxCoins([]int{2, 4, 1, 2, 7, 8}))
}
