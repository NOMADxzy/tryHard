package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}

	dp := make([][]bool, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ { // i,j 分别表示目前 截取的 字符串ab的长度
			if i == 0 && j == 0 {
				continue
			}
			k := i + j
			flag := false
			for l := 1; l <= i || l <= j; l++ {
				if i-l >= 0 && s1[i-l:i] == s3[k-l:k] && dp[i-l][j] {
					flag = true
					break
				}
				if j-l >= 0 && s2[j-l:j] == s3[k-l:k] && dp[i][j-l] {
					flag = true
					break
				}
			}
			dp[i][j] = flag
		}
	}
	return dp[m][n]
}

func main() {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
}
