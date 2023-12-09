package main

import "fmt"

func findNext(s string) []int {
	next := make([]int, len(s))
	i, l := 1, 0
	for i < len(s) {
		if s[i] == s[l] {
			l++
			next[i] = l
			i++
		} else {
			if l == 0 {
				i++
			} else {
				l = next[l-1]
			}
		}
	}
	return next
}

// KMP算法
func shortestPalindrome(s string) string {
	m := len(s)
	pre := make([]byte, m)
	for i := 0; i < m; i++ {
		pre[i] = s[m-1-i]
	}
	pattern := s + "#" + string(pre)
	next := findNext(pattern)

	preStr := string(pre[:m-next[len(next)-1]])
	return preStr + s
}

// 动规，超时
func shortestPalindrome1(s string) string {
	m := len(s)
	if m == 0 {
		return ""
	}
	dp := make([][]bool, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]bool, m)
		dp[i][i] = true
	}

	for l := 1; l < m; l++ {
		for i := 0; i < m-l; i++ {
			if s[i] == s[i+l] && (l == 1 || dp[i+1][i+l-1]) {
				dp[i][i+l] = true
			}
		}
	}
	var pre []byte
	for r := m - 1; r >= 0; r-- {
		if dp[0][r] {
			break
		}
		pre = append(pre, s[r])
	}
	return string(pre) + s
}

func main() {
	fmt.Println(shortestPalindrome("abcd"))
}
