package main

import "fmt"

type Trie struct {
	children [26]*Trie
	mark     bool
}

func insert(root *Trie, word string) {
	node := root
	for i := len(word) - 1; i >= 0; i-- {
		ch := word[i] - 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.mark = true
}

func search(root *Trie, word string) bool {
	node := root
	for i := len(word) - 1; i >= 0; i-- {
		ch := word[i] - 'a'
		if node.children[ch] == nil {
			return false
		}
		node = node.children[ch]
	}
	for _, child := range node.children {
		if child != nil {
			return true
		}
	}
	return false
}

func minimumLengthEncoding(words []string) int {
	dup := map[string]bool{}
	root := &Trie{}
	for _, word := range words {
		insert(root, word)
	}
	res := 0
	for _, word := range words {
		if !search(root, word) {
			if !dup[word] {
				dup[word] = true
				res += len(word) + 1
			}
		}
	}
	return res
}

func main() {
	fmt.Println(minimumLengthEncoding([]string{"time", "me", "bell"}))
}
