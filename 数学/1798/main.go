package main

import (
	"fmt"
	"sort"
)

func getMaximumConsecutive(coins []int) int {
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] < coins[j]
	})
	preMax := 0
	for i := 0; i < len(coins); i++ {
		cur := coins[i]
		if cur > preMax+1 { //不连续了
			return preMax + 1
		}
		newMax := preMax + cur
		preMax = newMax
	}
	return preMax + 1
}

func main() {
	fmt.Println(getMaximumConsecutive([]int{1, 1, 1, 4}))
}
