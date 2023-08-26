package main

import "fmt"

// dp中放字符串会爆内存
func shortestCommonSupersequence(str1 string, str2 string) string {
	m, n := len(str1), len(str2)

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}
	for i := 0; i < n+1; i++ {
		dp[0][i] = i
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				choice1 := dp[i-1][j] + 1
				choice2 := dp[i][j-1] + 1
				if choice1 < choice2 {
					dp[i][j] = choice1
				} else {
					dp[i][j] = choice2
				}
			}
		}
	}
	s := ""
	i, j := m, n
	for i > 0 && j > 0 {
		if str1[i-1] == str2[j-1] {
			s = str1[i-1:i] + s
			i--
			j--
		} else if dp[i-1][j] < dp[i][j-1] {
			s = str1[i-1:i] + s
			i--
		} else {
			s = str2[j-1:j] + s
			j--
		}
	}
	if i > 0 {
		s = str1[:i] + s
	}
	if j > 0 {
		s = str2[:j] + s
	}

	return s
}

func main() {
	fmt.Println(shortestCommonSupersequence("abac", "cab"))
}
