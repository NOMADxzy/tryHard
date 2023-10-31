package main

import (
	"fmt"
)

type Trie struct {
	Children [26]*Trie
	isEnd    bool
}

func findLongest(node *Trie, best *string, acc []byte) {
	for i := 0; i < 26; i++ {
		if node.Children[i] != nil && node.Children[i].isEnd {
			findLongest(node.Children[i], best, append(acc, byte('a'+i)))
		} else if len(*best) < len(acc) {
			*best = string(acc)
		}
	}
}

func longestWord(words []string) string {
	root := &Trie{}
	for _, word := range words {
		node := root
		for _, c := range word {
			if node.Children[c-'a'] == nil {
				node.Children[c-'a'] = &Trie{}
			}
			node = node.Children[c-'a']
		}
		node.isEnd = true
	}
	var acc []byte
	best := ""
	findLongest(root, &best, acc)
	return best
}

func main() {
	fmt.Println(longestWord([]string{"ab"}))
}
