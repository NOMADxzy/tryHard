package main

import "fmt"

func minimumDeleteSum(s1 string, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		dp[i] = make([]int, len(s2)+1)
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + int(s1[i-1]) + int(s2[j-1])
			} else {
				val := dp[i-1][j-1]
				var p int
				for p = j - 2; p >= 0 && s1[i-1] != s2[p]; p-- {
				}
				if p >= 0 {
					val1 := dp[i-1][p] + int(s1[i-1]) + int(s2[p])
					if val1 > val {
						val = val1
					}
				}
				for p = i - 2; p >= 0 && s1[p] != s2[j-1]; p-- {
				}
				if p >= 0 {
					val2 := dp[p][j-1] + int(s1[p]) + int(s2[j-1])
					if val2 > val {
						val = val2
					}
				}
				dp[i][j] = val
			}
		}
	}
	total := 0
	for i := 0; i < len(s2); i++ {
		total += int(s2[i])
	}
	for i := 0; i < len(s1); i++ {
		total += int(s1[i])
	}
	return total - dp[len(s1)][len(s2)]
}

func main() {
	fmt.Println(minimumDeleteSum("delete", "leet"))
}
