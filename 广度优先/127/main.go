package main

import "fmt"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordList = append(wordList, beginWord)
	m := len(wordList)
	wordMap := map[string]int{}
	for i := 0; i < len(wordList); i++ {
		wordMap[wordList[i]] = i
	}
	var queue []int
	queue = append(queue, m-1)
	visited := make([]bool, m)
	dist := 1
	tarIdx, ok := wordMap[endWord]
	if !ok {
		return 0
	}
	for len(queue) > 0 {
		var nextQ []int
		for len(queue) > 0 {
			wordIdx := queue[0]
			if wordIdx == tarIdx {
				return dist
			}
			word := make([]byte, len(wordList[wordIdx]))
			for i := 0; i < len(word); i++ {
				word[i] = wordList[wordIdx][i]
			}
			visited[wordIdx] = true
			queue = queue[1:]
			for i := 0; i < len(word); i++ {
				preChar := word[i]
				for j := 0; j < 26; j++ {
					word[i] = byte('a' + j)
					if idx, ok := wordMap[string(word)]; ok && !visited[idx] {
						nextQ = append(nextQ, idx)
					}
				}
				word[i] = preChar
			}
		}
		queue = nextQ
		dist++
	}
	return 0
}

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}
