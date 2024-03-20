package main

import "fmt"

func findAnagrams(s string, p string) []int {
	limit := map[byte]int{}
	for i := 0; i < len(p); i++ {
		limit[p[i]]++
	}
	l, r := 0, 0
	tmp := map[byte]int{}

	var ans []int
	for r < len(s) {
		tmp[s[r]]++
		for tmp[s[r]] > limit[s[r]] {
			tmp[s[l]]--
			l++
		}
		if r-l+1 == len(p) {
			ans = append(ans, l)
		}
		r++
	}
	return ans
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd0", "abc"))
}
