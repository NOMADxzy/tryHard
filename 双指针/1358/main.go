package main

import "fmt"

func numberOfSubstrings(s string) int {
	kMap := map[byte]int{}
	cnt := 0
	ans := 0
	l, r := 0, 0
	for r < len(s) {
		kMap[s[r]]++
		if kMap[s[r]] == 1 {
			cnt++
		}
		for cnt == 3 {
			ans += len(s) - r
			kMap[s[l]]--
			if kMap[s[l]] == 0 {
				cnt--
			}
			l++
		}
		r++
	}
	return ans
}

func main() {
	fmt.Println(numberOfSubstrings("abcabc"))
}
