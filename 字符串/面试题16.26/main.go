package main

import (
	"fmt"
	"strings"
)

func calculate(s string) int {
	var symbs []int
	s = strings.ReplaceAll(s, " ", "")
	idx := 0
	if s[0] == '-' {
		idx++
		symbs = append(symbs, -1)
	} else {
		symbs = append(symbs, 1)
	}
	var elem []string
	for idx < len(s) {
		nextIdx := idx
		for nextIdx < len(s) && s[nextIdx] != '+' && s[nextIdx] != '-' {
			nextIdx++
		}
		elem = append(elem, s[idx:nextIdx])
		idx = nextIdx
		if idx < len(s) {
			if s[idx] == '-' {
				symbs = append(symbs, -1)
			} else {
				symbs = append(symbs, 1)
			}
		}
		idx++
	}
	ans := 0
	var getVal func(s string) int
	getVal = func(s string) int {
		curVal := 0
		powVal := 1
		var i int
		for i = len(s) - 1; i >= 0 && s[i] >= '0' && s[i] <= '9'; i-- {
			curVal += powVal * int(s[i]-'0')
			powVal *= 10
		}
		if i == -1 {
			return curVal
		} else if s[i] == '*' {
			return getVal(s[:i]) * curVal
		} else {
			return getVal(s[:i]) / curVal
		}
	}
	for i := 0; i < len(symbs); i++ {
		ans += symbs[i] * getVal(elem[i])
	}
	return ans
}

func main() {
	fmt.Println(calculate("3+5 / 2"))
}
