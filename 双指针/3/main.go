package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	charMap := map[uint8]int{}
	l, r := 0, 0
	length := 0

	for r = 0; r < len(s); r++ {
		c := s[r]
		charMap[c] += 1
		if charMap[c] > 1 {
			for s[l] != c {
				charMap[s[l]] -= 1
				l++
			}
			l++
			charMap[c] -= 1
			if r-l+1 > length {
				length = r - l + 1
			}
		} else {
			if r-l+1 > length {
				length = r - l + 1
			}
		}
	}
	return length
}

func main() {
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
}
