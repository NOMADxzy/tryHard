package main

import (
	"fmt"
	"sort"
)

func longestStrChain(words []string) int {
	m := len(words)
	dp := make([]int, m)

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	dp[0] = 1
	ans := 1
	idxMap := map[string]int{words[0]: 0}

	for i := 1; i < m; i++ {
		idxMap[words[i]] = i
		dp[i] = 1
		for j := 0; j < len(words[i]); j++ {
			if idx, ok := idxMap[words[i][:j]+words[i][j+1:]]; ok {
				if dp[idx]+1 > dp[i] {
					dp[i] = dp[idx] + 1
				}
			}
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

func main() {
	fmt.Println(longestStrChain([]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}))
}
