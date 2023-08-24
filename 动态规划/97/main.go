package main

import (
	"fmt"
)

func isInterleaveSingle(s1 string, s2 string, s3 string) bool {
	if s2 == "" {
		return s3 == s1
	}
	if len(s3) != len(s1)+len(s2) {
		return false
	}
	for i := 1; i <= len(s1); i++ {
		if s3[:i] == s1[:i] {

			if isInterleave(s2, s1[i:], s3[i:]) {
				return true
			}
		} else {
			break
		}
	}
	return false
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	//return isInterleaveSingle(s1, s2, s3) || isInterleaveSingle(s2, s1, s3)
	//return solution(s1, s2, s3) || solution(s2, s1, s3)
	return dp(s1, s2, s3)
}

func solution(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	} else if s1 == "" {
		return s3 == s2
	} else if s2 == "" {
		return s3 == s1
	} else {
		if s1[0] == s3[0] {
			if solution(s1[1:], s2, s3[1:]) {
				return true
			}
		} else {
			if s2[0] != s3[0] {
				return false
			}
			if solution(s2, s1, s3) {
				return true
			}
		}
		return false
	}
}

func dp(s1 string, s2 string, s3 string) bool {
	m := len(s1)
	n := len(s2)
	if len(s3) != m+n {
		return false
	}

	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true

	for i := 1; i <= m; i++ {
		dp[i][0] = s1[:i] == s3[:i]
	}
	for i := 1; i <= n; i++ {
		dp[0][i] = s2[:i] == s3[:i]
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = dp[i-1][j] && s3[i+j-1] == s1[i-1] || dp[i][j-1] && s3[i+j-1] == s2[j-1]
		}

	}
	return dp[m][n]

}

func main() {
	//fmt.Println(isInterleave("bcbccabcccbcbbbcbbacaaccccacbaccabaccbabccbabcaabbbccbbbaa", "ccbccaaccabacaabccaaccbabcbbaacacaccaacbacbbccccbac", "bccbcccabbccaccaccacbacbacbabbcbccbaaccbbaacbcbaacbacbaccaaccabcaccacaacbacbacccbbabcccccbababcaabcbbcccbbbaa"))
	fmt.Println(isInterleave("a", "", ""))
}
