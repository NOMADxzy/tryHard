package main

import "fmt"

func wordSubsets(words1 []string, words2 []string) []string {
	word2MaxMap := map[byte]int{}
	tmpMap := map[byte]int{}
	for _, word := range words2 {
		for i := 0; i < len(word); i++ {
			tmpMap[word[i]]++
		}
		for b, cnt := range tmpMap {
			if cnt > word2MaxMap[b] {
				word2MaxMap[b] = cnt
			}
		}
		clear(tmpMap)
	}
	var ans []string
loop:
	for _, word := range words1 {
		clear(tmpMap)
		for i := 0; i < len(word); i++ {
			tmpMap[word[i]]++
		}
		for b, cnt := range word2MaxMap {
			if cnt > tmpMap[b] {
				continue loop
			}
		}
		ans = append(ans, word)
	}
	return ans
}

func main() {
	fmt.Println(wordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"lo", "eo"}))
}
