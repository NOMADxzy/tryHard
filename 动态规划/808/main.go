package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func soupServings(n int) float64 {

	cnt := (n + 24) / 25
	if cnt >= 179 {
		return 1
	}
	//if n%25 != 0 {
	//	cnt++
	//}
	dp := make([][]float64, cnt+1)
	for i := 0; i <= cnt; i++ {
		dp[i] = make([]float64, cnt+1)
	}

	dp[0][0] = 0.5
	for i := 1; i <= cnt; i++ {
		dp[0][i] = 1
	}

	for i := 1; i <= cnt; i++ {
		for j := 1; j <= cnt; j++ {
			p := dp[max(i-4, 0)][j] + dp[max(i-3, 0)][max(j-1, 0)] + dp[max(i-2, 0)][max(j-2, 0)] + dp[max(i-1, 0)][max(j-3, 0)]
			dp[i][j] = p / 4
		}
	}
	return dp[cnt][cnt]
}

func main() {
	fmt.Println(soupServings(0))
}
