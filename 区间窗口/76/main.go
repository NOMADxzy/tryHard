package main

import "fmt"

func minWindow(s string, t string) string {
	tMap := map[byte]int{}
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}
	left, right := 0, len(s)-1
	sMap := map[byte]int{}
	for i := 0; i < len(s); i++ {
		sMap[s[i]]++
	}
	for key, val := range tMap {
		if sMap[key] < val {
			return ""
		}
	}
	for sMap[s[right]] > tMap[s[right]] {
		sMap[s[right]]--
		right--
	}
	sMap[s[right]]--
	ans := [2]int{left, right}
	for ; right < len(s); right++ {
		sMap[s[right]]++
		for sMap[s[left]] > tMap[s[left]] {
			sMap[s[left]]--
			left++
		}
		if right-left < ans[1]-ans[0] {
			ans[0] = left
			ans[1] = right
		}
	}
	return s[ans[0] : ans[1]+1]
}

func main() {
	fmt.Println(minWindow("cabwefgewcwaefgcf", "cae"))
}
