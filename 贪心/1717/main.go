package main

import "fmt"

func maximumGain(s string, x int, y int) int {
	a, b := byte('a'), byte('b')
	if x < y {
		a, b = b, a
		x, y = y, x
	}
	ans := 0
	cntb, cnta := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] != a && s[i] != b {
			ans += min(cnta, cntb) * y
			cntb = 0
			cnta = 0
		} else if s[i] == a {
			cnta++
		} else {
			if cnta > 0 {
				cnta--
				ans += x
			} else {
				cntb++
			}
		}
	}
	ans += min(cnta, cntb) * y
	return ans
}

func main() {
	fmt.Println(maximumGain("cdbcbbaaabab", 4, 5))
}
