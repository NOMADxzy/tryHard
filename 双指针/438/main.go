package main

import "fmt"

func findAnagrams(s string, p string) []int {
	pMap := map[byte]int{}
	for i := 0; i < len(p); i++ {
		pMap[p[i]]++
	}

	l, r := 0, 0
	var idxs []int
	tmpMap := map[byte]int{}

	for ; r < len(s); r++ {
		tmpMap[s[r]]++
		if tmpMap[s[r]] > pMap[s[r]] {
			for s[l] != s[r] {
				tmpMap[s[l]]--
				l++
			}
			tmpMap[s[l]]--
			l++
		}
		if r-l+1 == len(p) {
			idxs = append(idxs, l)
			tmpMap[s[l]]--
			l++
		}
	}
	return idxs
}

func main() {
	fmt.Println(findAnagrams("abab", "ab"))
}
