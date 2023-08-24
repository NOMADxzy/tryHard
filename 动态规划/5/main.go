package main

import "fmt"

// TODO 注意不要漏掉长度为1的字符串
func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i, _ := range dp {
		dp[i] = make([]bool, len(s))
	}
	min := 0
	cur := ""

	for l := 1; l <= len(s); l++ {
		for i := 0; i <= len(s)-l; i++ {
			if l == 1 {
				dp[i][i] = true
				if l > min {
					min = l
					cur = s[i : i+l]
				}
			} else if l >= 2 {
				if s[i] != s[i+l-1] {
					dp[i][i+l-1] = false
				} else {
					dp[i][i+l-1] = l == 2 || dp[i+1][i+l-2]
					if dp[i][i+l-1] && l > min {
						min = l
						cur = s[i : i+l]
					}
				}
			}
		}
	}
	return cur
}
func main() {
	s := "ac"
	fmt.Println(longestPalindrome(s))
}
