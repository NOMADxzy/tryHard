package main

import "fmt"

func longestSemiRepetitiveSubstring(s string) int {
	nextPlace := -1
	var ans int

	l, r := 0, 0
	for r < len(s) {
		if r > 0 && s[r] == s[r-1] {
			if nextPlace > 0 {
				l = nextPlace
			}
			nextPlace = r
		}
		ans = max(ans, r-l+1)
		r++
	}
	return ans
}

func main() {
	fmt.Println(longestSemiRepetitiveSubstring("52233"))
}
