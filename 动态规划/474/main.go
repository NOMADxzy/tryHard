package main

import "fmt"

func count01(str string) (int, int) {
	cnt0, cnt1 := 0, 0
	for i := 0; i < len(str); i++ {
		if str[i] == '0' {
			cnt0++
		} else {
			cnt1++
		}
	}
	return cnt0, cnt1
}

func findMaxForm(strs []string, m int, n int) int {
	K := len(strs)
	dp := make([][][]int, K+1)
	for i := 0; i <= K; i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}

	for k := 1; k <= K; k++ {
		str := strs[k-1]
		cnt0, cnt1 := count01(str)
		for i := 0; i <= m; i++ {
			for j := 0; j <= n; j++ {
				dp[k][i][j] = dp[k-1][i][j]
				if i-cnt0 >= 0 && j-cnt1 >= 0 {
					dp[k][i][j] = max(dp[k][i][j], dp[k-1][i-cnt0][j-cnt1]+1)
				}
			}
		}
	}
	return dp[K][m][n]
}

func main() {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
}
