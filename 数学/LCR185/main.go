package main

import "fmt"

func statisticsProbability(num int) []float64 {
	dp := make([][]float64, num)
	dp[0] = []float64{float64(1) / float64(6), float64(1) / float64(6), float64(1) / float64(6), float64(1) / float64(6), float64(1) / float64(6), float64(1) / float64(6)}
	for i := 1; i < num; i++ {
		dp[i] = make([]float64, 5*(i+1)+1)
		lastStart, lastEnd := i, 6*i
		for sum := i + 1; sum <= 6*(i+1); sum++ {
			for cur := 1; cur <= 6; cur++ {
				remain := sum - cur
				if remain >= lastStart && remain <= lastEnd {
					dp[i][sum-(i+1)] += dp[i-1][remain-lastStart] * float64(1) / float64(6)
				}
			}
		}
	}
	return dp[num-1]
}

func main() {
	fmt.Println(statisticsProbability(3))
}
