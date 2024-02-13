package main

import (
	"fmt"
	"sort"
)

func makeLargestSpecial(s string) string {
	var findBest func(s string) string
	findBest = func(s string) string {
		if len(s) <= 2 {
			return s
		}
		var subStrings []string
		cnt1 := 0
		l := 0
		for i := 0; i < len(s); i++ {
			if s[i] == '1' {
				cnt1++
			} else {
				cnt1--
				if cnt1 == 0 {
					subStrings = append(subStrings, "1"+findBest(s[l+1:i])+"0")
					l = i + 1
				}
			}
		}
		ans := "" // 10 101100
		sort.Slice(subStrings, func(i, j int) bool {
			return subStrings[i] > subStrings[j]
		})
		for _, subString := range subStrings {
			ans += subString
		}
		return ans
	}
	return findBest(s)
}

func main() {
	fmt.Println(makeLargestSpecial("11011000"))
}
