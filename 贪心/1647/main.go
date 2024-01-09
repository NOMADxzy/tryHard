package main

import (
	"fmt"
	"sort"
)

func minDeletions(s string) int {
	kMap := map[byte]int{}
	for i := 0; i < len(s); i++ {
		kMap[s[i]]++
	}
	var cnts []int
	for _, cnt := range kMap {
		cnts = append(cnts, cnt)
	}
	sort.Slice(cnts, func(i, j int) bool {
		return cnts[i] > cnts[j]
	})
	ans := 0
	pre := 0
	for i := 1; i < len(cnts); i++ {
		for cnts[i] >= cnts[pre] {
			ans++
			cnts[i]--
		}
		if cnts[i] != 0 {
			pre = i
		}
	}
	return ans
}

func main() {
	fmt.Println(minDeletions("aaabbbccdd"))
}
