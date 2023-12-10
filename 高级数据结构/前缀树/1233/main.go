package main

import (
	"fmt"
	"strings"
)

type Trie struct {
	Children map[string]*Trie
	ref      int
}

func dfs(root *Trie, ans *[]string, folder []string) {
	if root.ref >= 0 {
		*ans = append(*ans, folder[root.ref])
		return
	}
	for _, node := range root.Children {
		dfs(node, ans, folder)
	}
}

func removeSubfolders(folder []string) []string {
	root := &Trie{Children: map[string]*Trie{}, ref: -1}
	for i, s := range folder {
		path := strings.Split(s[1:], "/")
		node := root
		for _, f := range path {
			if node.Children[f] == nil {
				node.Children[f] = &Trie{Children: map[string]*Trie{}, ref: -1}
			}
			node = node.Children[f]
		}
		node.ref = i
	}
	ans := make([]string, 0, 2*len(folder))
	dfs(root, &ans, folder)
	return ans
}

func main() {
	folder := []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}
	fmt.Println(removeSubfolders(folder))
}
