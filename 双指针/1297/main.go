package main

import "fmt"

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	kMap := map[byte]int{}
	sMap := map[string]int{}
	cnt := 0
	ans := 0
	l, r := 0, 0
	for r < len(s) {
		kMap[s[r]]++
		if kMap[s[r]] == 1 {
			cnt++
		}
		for cnt > maxLetters || r-l+1 > minSize {
			kMap[s[l]]--
			if kMap[s[l]] == 0 {
				cnt--
			}
			l++
		}
		cur := r - l + 1
		if cur == minSize {
			sMap[s[l:r+1]]++
		}
		r++
	}
	for _, c := range sMap {
		if c > ans {
			ans = c
		}
	}
	return ans
}

func main() {
	fmt.Println(maxFreq("aabcabcab", 2, 1, 3))
}
