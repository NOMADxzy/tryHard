package main

import "fmt"

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		nextDp := make([]int, amount+1)
		copy(nextDp, dp)
		for a := 0; a <= amount; a++ {
			if dp[a] > 0 {
				for times := 1; times*coin+a <= amount; times++ {
					nextDp[times*coin+a] += dp[a]
				}
			}
		}
		dp = nextDp
	}
	return dp[amount]
}

func main() {
	fmt.Println(change(500, []int{1, 2, 5}))
}
