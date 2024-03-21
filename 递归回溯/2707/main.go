package main

import "fmt"

func minExtraChar(s string, dictionary []string) int {
	hist := make([]int, len(s))
	for i := 0; i < len(hist); i++ {
		hist[i] = -1
	}
	var dfs func(i int) int
	dfs = func(i int) int {
		if i == len(s) {
			return 0
		}
		if hist[i] >= 0 {
			return hist[i]
		}
		minVal := len(s) - i
		for j := 0; j < len(dictionary); j++ {
			word := dictionary[j]
			if i+len(word) <= len(s) && s[i:i+len(word)] == word {
				minVal = min(minVal, dfs(i+len(word)))
			}
		}
		minVal = min(minVal, 1+dfs(i+1)) // 不匹配任何word，直接放弃一位
		hist[i] = minVal
		return minVal
	}
	return dfs(0)
}

func main() {
	fmt.Println(minExtraChar("leeteeeee", []string{"leet", "teeeee", "code", "leetcode"}))
}
