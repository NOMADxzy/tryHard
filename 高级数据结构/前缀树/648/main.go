package main

import (
	"fmt"
	"strings"
)

type Trie struct {
	Children [26]*Trie
	end      bool
}

func replaceWords(dictionary []string, sentence string) string {
	root := &Trie{}
	for _, s := range dictionary {
		node := root
		for i := 0; i < len(s); i++ {
			idx := s[i] - 'a'
			if node.Children[idx] == nil {
				node.Children[idx] = &Trie{}
			}
			node = node.Children[idx]
		}
		node.end = true
	}
	find := func(s string) int {
		node := root
		var i int
		for i = 0; i < len(s); i++ {
			idx := s[i] - 'a'
			if node.Children[idx] == nil || node.end {
				break
			}
			node = node.Children[idx]
		}
		if node.end {
			return i
		} else {
			return len(s)
		}
	}
	sentSplits := strings.Fields(sentence)
	for i, s := range sentSplits {
		sentSplits[i] = s[:find(s)]
	}
	return strings.Join(sentSplits, " ")
}

func main() {
	fmt.Println(replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"))
}
