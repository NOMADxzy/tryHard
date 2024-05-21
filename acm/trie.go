package main

import "fmt"

type Trie struct {
	Children [26]*Trie
	Mark     bool
	Val      string
}

func buildTrie(words []string) *Trie {
	root := &Trie{}
	for _, word := range words {
		cur := root
		for i := 0; i < len(word); i++ {
			bIdx := int(word[i] - 'a')
			if cur.Children[bIdx] == nil {
				cur.Children[bIdx] = &Trie{}
			}
			cur = cur.Children[bIdx]
		}
		cur.Mark = true
		cur.Val = word
	}
	return root
}

func Query(root *Trie, target string) []string {
	var ans []string
	cur := root
	for i := 0; i < len(target); i++ {
		b := target[i]
		bIdx := int(b - 'a')
		if cur.Children[bIdx] == nil {
			return []string{}
		}
		cur = cur.Children[bIdx]
	}
	var dfs func(root *Trie)
	dfs = func(root *Trie) {
		if root == nil {
			return
		}
		if root.Mark {
			ans = append(ans, root.Val)
		}
		for _, child := range root.Children {
			dfs(child)
		}
	}
	dfs(cur)
	return ans
}

func main() {
	words := []string{"abcd", "aba", "bcd"}
	root := buildTrie(words)
	fmt.Println(Query(root, "a"))
	fmt.Println(Query(root, "abc"))
	fmt.Println(Query(root, "b"))
	fmt.Println(Query(root, "bcd"))
	fmt.Println(Query(root, "e"))

}
