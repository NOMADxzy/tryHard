package main

import "fmt"

func minimumCoins(prices []int) int {
	m := len(prices)
	dp := make([]int, m+1)

	for i := 0; i < m; i++ {
		minVal := dp[i+1-1] + prices[i]
		for j := i - 1; j >= 0; j-- {
			if 2*(j+1)-1 >= i {
				if v := dp[j] + prices[j]; v < minVal {
					minVal = v
				}
			}
		}
		dp[i+1] = minVal
	}
	return dp[m]
}

func main() {
	fmt.Println(minimumCoins([]int{3, 1, 2}))
}
