package main

import "fmt"

func twoEggDrop(n int) int {
	dp := make([]int, n+1)

	dp[1] = 1
	//dp[2] = 2

	for i := 2; i <= n; i++ {
		var j int
		for j = 1; j <= 1+dp[i-j]; j++ {
		}
		j--
		dp[i] = 1 + dp[i-j]
	}
	return dp[n]
}

func main() {
	fmt.Println(twoEggDrop(100))
}
