package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	cnt := map[byte]int{}
	for i := 0; i < len(s1); i++ {
		cnt[s1[i]]--
	}

	l, r := 0, 0
	for r < len(s2) {
		cnt[s2[r]]++
		for cnt[s2[r]] > 0 {
			cnt[s2[l]]--
			l++
		}
		if r-l+1 == len(s1) { // 原理：a<=2, b<=2 & a + b = 4 则一定有 a==2, b==2
			return true
		}
		r++
	}
	return false
}

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
}
