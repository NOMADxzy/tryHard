package main

import (
	"fmt"
	"sort"
)

func characterReplacement(s string, k int) int {
	charCnts := map[byte]int{}
	var chars []byte

	for i := 0; i < len(s); i++ {
		charCnts[s[i]]++
		if charCnts[s[i]] == 1 {
			chars = append(chars, s[i])
		}
	}

	sort.Slice(chars, func(i, j int) bool {
		return charCnts[chars[i]] > charCnts[chars[j]]
	})

	maxLen := 0
	for i := 0; i < len(chars); i++ {
		c := chars[i]
		if charCnts[c]+k <= maxLen {
			break
		}
		cnt := 0
		l, r := 0, 0
		for ; r < len(s); r++ {
			if s[r] == c {
				cnt++
			}
			for r-l+1 > k+cnt && r < len(s) {
				if s[l] == c {
					cnt--
				}
				l++
				r++
				if r < len(s) && s[r] == c {
					cnt++
				}
			}
			if r == len(s) {
				break
			}
			if r-l+1 > maxLen {
				maxLen = r - l + 1
			}
		}
	}

	return maxLen
}

func main() {
	fmt.Println(characterReplacement("AAAA", 2))
}
