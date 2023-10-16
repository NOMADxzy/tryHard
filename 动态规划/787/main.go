package main

import "fmt"

//优化后的二维动态规划 勉强过，建议使用贝尔曼方法
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	dp := make([][][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, n)
		}
	}
	for _, f := range flights {
		dp[0][f[0]][f[1]] = f[2]
	}

	for K := 1; K <= k; K++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				minCost := 1000000000
				for c := 0; c < n; c++ {
					if c != i && c != j && dp[K-1][i][c] > 0 && dp[0][c][j] > 0 {
						cost := dp[K-1][i][c] + dp[0][c][j]
						if cost < minCost {
							minCost = cost
						}
					}
				}
				dp[K][i][j] = dp[K-1][i][j]
				if dp[K][i][j] == 0 || minCost < dp[K][i][j] {
					dp[K][i][j] = minCost
				}
			}
		}
	}
	if dp[k][src][dst] == 0 || dp[k][src][dst] == 1000000000 {
		return -1
	} else {
		return dp[k][src][dst]
	}
}

func main() {
	flights := [][]int{{0, 1, 2}, {1, 2, 1}, {2, 0, 10}}
	fmt.Println(findCheapestPrice(3, flights, 1, 2, 1))
}
