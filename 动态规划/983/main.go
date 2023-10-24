package main

import "fmt"

func mincostTickets(days []int, costs []int) int {
	m := len(days)
	dp := make([]int, m)
	minCost := costs[0]
	for i := 1; i < len(costs); i++ {
		if costs[i] < minCost {
			minCost = costs[i]
		}
	}
	dp[0] = minCost

	for i := 1; i < m; i++ {
		best := minCost + dp[i-1]
		for j := i - 1; j >= 0; j-- {
			if days[i]-days[j] < 7 {
				preCost := 0
				if j > 0 {
					preCost = dp[j-1]
				}
				if costs[1]+preCost < best {
					best = costs[1] + preCost
				}
			}
			if days[i]-days[j] < 30 {
				preCost := 0
				if j > 0 {
					preCost = dp[j-1]
				}
				if costs[2]+preCost < best {
					best = costs[2] + preCost
				}
			} else if days[i]-days[j] >= 30 {
				break
			}
		}
		dp[i] = best
	}
	return dp[m-1]
}

func main() {
	fmt.Println(mincostTickets([]int{1, 7}, []int{3, 2, 1}))
}
