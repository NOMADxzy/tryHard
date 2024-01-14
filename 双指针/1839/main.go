package main

import "fmt"

func longestBeautifulSubstring(word string) int {
	idxMap := map[byte]int{'a': 0, 'e': 1, 'i': 2, 'o': 3, 'u': 4}
	//cnts := make([]int, 5)
	//lastIdx := make([]int, 5)

	l, r := 0, 1
	//cnts[idxMap[word[l]]]++
	//lastIdx[idxMap[word[l]]] = 1
	ans := 0
	for r < len(word) {
		if !(word[r] == word[r-1] || idxMap[word[r]] == idxMap[word[r-1]]+1) || word[l] != 'a' {
			l = r
		}
		if word[l] == 'a' && word[r] == 'u' {
			ans = max(ans, r-l+1)
			//lastIdx[idxMap[word[r]]] = r
		}
		r++
	}
	return ans
}

func main() {
	fmt.Println(longestBeautifulSubstring("aeeeiiiioooauuuaeiou"))
}
