package main

import "fmt"

func findSubstring(s string, words []string) []int {
	delta := len(words[0])
	tarLen := delta * len(words)
	var ans []int
	wordsMap := map[string]int{}
	for i := 0; i < len(words); i++ {
		wordsMap[words[i]]++
	}
	tmpMap := map[string]int{}
	for L := 0; L < delta; L++ {
		clear(tmpMap)
		l := L
		for i := l + delta; i <= len(s); i += delta {
			cur := s[i-delta : i]
			tmpMap[cur]++
			for tmpMap[cur] > wordsMap[cur] {
				l += delta
				tmpMap[s[l-delta:l]]--
			}
			if i-l == tarLen {
				ans = append(ans, l)
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(findSubstring("lingmindraboofooowingdingbarrwingmonkeypoundcake", []string{"fooo", "barr", "wing", "ding", "wing"}))
}
