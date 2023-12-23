package main

import "fmt"

type Trie struct {
	Children [26]*Trie
	mark     bool
}

func minimumLengthEncoding(words []string) int {
	root := &Trie{}
	for _, word := range words {
		node := root
		for i := len(word) - 1; i >= 0; i-- {
			idx := word[i] - 'a'
			if node.Children[idx] == nil {
				node.Children[idx] = &Trie{}
			}
			node = node.Children[idx]
		}
		node.mark = true
	}
	var dfs func(root *Trie, acc int) (int, int)
	dfs = func(root *Trie, acc int) (int, int) {
		leaf := true
		length := 0
		cnt := 0
		for _, child := range root.Children {
			if child != nil {
				leaf = false
				c, l := dfs(child, acc+1)
				cnt += c
				length += l
			}
		}
		if leaf {
			return 1, acc
		}
		return cnt, length
	}
	cnt, length := dfs(root, 0)
	ans := cnt + length
	return ans
}

func main() {
	fmt.Println(minimumLengthEncoding([]string{"time", "me", "bell"}))
}
