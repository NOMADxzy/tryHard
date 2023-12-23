package main

import "fmt"

func minDeletionSize(strs []string) int {
	n := len(strs[0])
	dp := make([]int, n)
	dp[0] = 1
	maxLen := 1
	for i := 1; i < n; i++ {
		best := 0
		for j := i - 1; j >= 0; j-- {
			ok := true
			for _, str := range strs {
				if str[i] < str[j] {
					ok = false
					break
				}
			}
			if ok && dp[j] > best {
				best = dp[j]
			}
		}
		dp[i] = best + 1
		maxLen = max(maxLen, dp[i])
	}
	return n - maxLen
}

func main() {
	fmt.Println(minDeletionSize([]string{"ghi", "def", "abc"}))
}
