package main

func LongestCommonSubsequence(str1 string, str2 string) string {
	m, n := len(str1), len(str2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	ans := ""
	ansLen := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
				if dp[i][j] > ansLen {
					ansLen = dp[i][j]
					ans = str2[j-ansLen : j]
				}
			}
		}
	}
	return ans
}
