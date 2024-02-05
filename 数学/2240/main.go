package main

import "fmt"

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	dp := make([]int64, total+1)
	items := []int{cost1, cost2}
	for _, item := range items { // dp[i]表示 有i数量钱时购买任意数量前k个item的方案数
		for i := item; i <= total; i++ {
			dp[i] += dp[i-item] + 1
		}
	}
	return dp[total] + 1
}

func main() {
	fmt.Println(waysToBuyPensPencils(4, 2, 1))
}
