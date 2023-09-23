package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	sMap := map[byte]byte{}
	used := map[byte]bool{}

	if len(s) != len(t) {
		return false
	}
	n := len(s)

	for i := 0; i < n; i++ {
		l := s[i]
		r := t[i]
		if sMap[l] > 0 {
			if sMap[l] != r {
				return false
			}
		} else if !used[r] {
			sMap[l] = r
			used[r] = true
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("paper", "title"))
}
