package main

import "fmt"

type Trie struct {
	Children [26]*Trie
	mark     bool
}

func findNextWord(root *Trie, s string) []string {
	if s == "" {
		return []string{""}
	}
	var ans []string
	node := root
	pos := 0
	for pos < len(s) && node.Children[s[pos]-'a'] != nil {
		node = node.Children[s[pos]-'a']
		if node.mark {
			nextPaths := findNextWord(root, s[pos+1:])
			for _, path := range nextPaths {
				if path != "" {
					ans = append(ans, s[:pos+1]+" "+path)
				} else {
					ans = append(ans, s[:pos+1])
				}
			}
		}
		pos++
	}
	return ans
}

func wordBreak(s string, wordDict []string) []string {
	root := &Trie{}
	for _, word := range wordDict {
		node := root
		for _, c := range word {
			idx := c - 'a'
			if node.Children[idx] == nil {
				node.Children[idx] = &Trie{}
			}
			node = node.Children[idx]
		}
		node.mark = true
	}
	return findNextWord(root, s)
}

func wordBreak1(s string, wordDict []string) []string {
	m := len(s)
	word_max_len := 0
	wordMap := map[string]bool{}
	for _, w := range wordDict {
		word_max_len = max(word_max_len, len(w))
		wordMap[w] = true
	}
	dp := make([][]string, m+1)
	dp[0] = append(dp[0], "")
	for i := 1; i <= m; i++ {
		for j := i - 1; j >= 0 && j+word_max_len > i-1; j-- {
			if wordMap[s[j:i]] {
				for _, preStr := range dp[j] {
					if preStr == "" {
						dp[i] = append(dp[i], s[j:i])
					} else {
						dp[i] = append(dp[i], preStr+" "+s[j:i])
					}
				}
			}
		}
	}
	return dp[m]
}

func main() {
	fmt.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
}
