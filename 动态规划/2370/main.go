package main

import "fmt"

func longestIdealString(s string, k int) int {
	m := len(s)
	lastPlace := make([]int, 26)
	for i := 0; i < 26; i++ {
		lastPlace[i] = -1
	}
	ans := 0
	dp := make([]int, m)
	for i := 0; i < m; i++ {
		dp[i] = 1
		for d := -k; d <= k; d++ {
			ch := int(s[i] - 'a')
			if ch+d >= 0 && ch+d < 26 && lastPlace[ch+d] >= 0 {
				dp[i] = max(dp[i], dp[lastPlace[ch+d]]+1)
			}
		}
		lastPlace[int(s[i]-'a')] = i
		ans = max(ans, dp[i])
	}
	return ans
}

func main() {
	fmt.Println(longestIdealString("aaaabd", 0))
}
