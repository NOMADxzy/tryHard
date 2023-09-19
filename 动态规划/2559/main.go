package main

import "fmt"

func isVowel(s string) bool {
	if (s[0] == 'a' || s[0] == 'e' || s[0] == 'i' || s[0] == 'o' || s[0] == 'u') &&
		(s[len(s)-1] == 'a' || s[len(s)-1] == 'e' || s[len(s)-1] == 'i' || s[len(s)-1] == 'o' || s[len(s)-1] == 'u') {
		return true
	}
	return false
}
func vowelStrings(words []string, queries [][]int) []int {
	dp := make([]int, len(words))

	if isVowel(words[0]) {
		dp[0] = 1
	}

	for i := 1; i < len(words); i++ {
		dp[i] = dp[i-1]
		if isVowel(words[i]) {
			dp[i]++
		}
	}

	res := make([]int, len(queries))
	for i := 0; i < len(res); i++ {
		if queries[i][0] == 0 {
			res[i] = dp[queries[i][1]]
		} else {
			res[i] = dp[queries[i][1]] - dp[queries[i][0]-1]
		}
	}
	return res
}

func main() {
	fmt.Println(vowelStrings([]string{"aba", "e", "ece"}, [][]int{{0, 2}, {0, 1}, {2, 2}}))
}
