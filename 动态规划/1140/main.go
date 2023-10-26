package main

import "fmt"

func stoneGameII(piles []int) int {
	m := len(piles)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m+1-i)
	}
	sums := make([]int, m)
	sums[0] = piles[0]
	for i := 1; i < m; i++ {
		sums[i] = sums[i-1] + piles[i]
	}

	dp[m-1][1] = piles[m-1]

	for i := m - 2; i >= 0; i-- {
		curDp := dp[i]
		for M := 1; M <= len(curDp)-1; M++ {
			best := 0
			for c := 1; c <= 2*M && c < m+1-i; c++ {
				nextM := M
				if c > nextM {
					nextM = c
				}
				theNextOneBest := 0
				if i+c < m {
					if nextM >= len(dp[i+c]) {
						nextM = len(dp[i+c]) - 1
					}
					theNextOneBest = dp[i+c][nextM]
				}
				R := sums[i+c-1]
				if i > 0 {
					R -= sums[i-1]
				}
				V := sums[m-1] - sums[i+c-1] - theNextOneBest
				if R+V > best {
					best = R + V
				}
			}
			dp[i][M] = best
		}
	}
	return dp[0][1]
}

func main() {
	fmt.Println(stoneGameII([]int{1, 2, 3, 4, 5, 100}))
}
