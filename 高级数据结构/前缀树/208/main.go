package main

type Trie struct {
	Children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, c := range word {
		id := c - 'a'
		if node.Children[id] == nil {
			node.Children[id] = &Trie{}
		}
		node = node.Children[id]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for _, c := range word {
		id := c - 'a'
		if node.Children[id] == nil {
			return false
		}
		node = node.Children[id]
	}
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, c := range prefix {
		id := c - 'a'
		if node.Children[id] == nil {
			return false
		}
		node = node.Children[id]
	}
	return true
}
