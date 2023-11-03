package main

import "fmt"

type Trie struct {
	Children [26]*Trie
	Mark     bool
}

func find3(res *[]string, root *Trie, acc string) {
	if len(*res) >= 3 {
		return
	}
	if root.Mark {
		*res = append(*res, acc)
	}
	for i := 0; i < 26; i++ {
		if root.Children[i] != nil {
			find3(res, root.Children[i], acc+string('a'+i))
		}
	}
}

func suggestedProducts(products []string, searchWord string) [][]string {
	root := &Trie{}
	for _, product := range products {
		node := root
		for _, c := range product {
			id := c - 'a'
			if node.Children[id] == nil {
				node.Children[id] = &Trie{}
			}
			node = node.Children[id]
		}
		node.Mark = true
	}

	m := len(searchWord)
	ans := make([][]string, m)
	node := root
	for i := 0; i < m; i++ {
		if node.Children[searchWord[i]-'a'] != nil {
			node = node.Children[searchWord[i]-'a']
			//var res []string
			find3(&ans[i], node, searchWord[:i+1])
		} else {
			break
		}
	}
	return ans
}

func main() {
	products := []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
	fmt.Println(suggestedProducts(products, "mo"))
}
