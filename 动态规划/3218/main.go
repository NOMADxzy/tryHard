package main

import "fmt"

func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int {
	dp := map[int]int{}
	var dfs func(startX, endX, startY, endY int) int
	dfs = func(startX, endX, startY, endY int) int {
		key := (n*startX+endX)*(m*m+m) + m*startY + endY
		if v, ok := dp[key]; ok {
			return v
		}
		minCost := 1 << 31
		for i := startY; i < endY; i++ {
			cost := horizontalCut[i] + dfs(startX, endX, startY, i) + dfs(startX, endX, i+1, endY)
			if cost < minCost {
				minCost = cost
			}
		}
		for j := startX; j < endX; j++ {
			cost := verticalCut[j] + dfs(startX, j, startY, endY) + dfs(j+1, endX, startY, endY)
			if cost < minCost {
				minCost = cost
			}
		}
		if minCost == 1<<31 {
			minCost = 0
		}
		dp[key] = minCost
		return minCost
	}
	return dfs(0, n-1, 0, m-1)
}

func main() {
	fmt.Println(minimumCost(3, 2, []int{1, 3}, []int{5}))
}
