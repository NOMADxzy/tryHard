package main

import "fmt"

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func maxVowels(s string, k int) int {
	m := len(s)
	l, r := 0, 0
	cnt := 0
	max := 0

	for ; r < m; r++ {
		if isVowel(s[r]) {
			cnt++
		}
		if r-l+1 > k {
			if isVowel(s[l]) {
				cnt--
			}
			l++
		}
		if cnt > max {
			max = cnt
		}
	}
	return max
}

func main() {
	fmt.Println(maxVowels("abciiidef", 3))
}
