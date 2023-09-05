package main

import (
	"fmt"
	"sort"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	sort.Slice(coins, func(i, j int) bool {
		return coins[i] < coins[j] //注意此处排序
	})

	for i := 1; i <= amount; i++ {
		min := -1
		for j := 0; j < len(coins) && coins[j] <= i; j++ {
			val := dp[i-coins[j]]
			if val >= 0 {
				if min == -1 {
					min = 1 + val
				} else if 1+val < min {
					min = 1 + val
				}
			}
		}
		dp[i] = min
	}
	return dp[amount]
}

func main() {
	fmt.Println(coinChange([]int{474, 83, 404, 3}, 264))
}
