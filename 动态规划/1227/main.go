package main

import "fmt"

func nthPersonGetsNthSeat(n int) float64 {
	if n == 1 {
		return 1
	}

	dp := make([]float64, n)
	dp[n-1] = float64(0)

	for i := n - 2; i >= 0; i-- {
		curRate := float64(1)
		curRate += dp[i+1] * float64(n-i-2)
		curRate /= float64(n - i)
		dp[i] = curRate
	}
	return dp[0]
}

func main() {
	fmt.Println(nthPersonGetsNthSeat(2))
}
