package main

import "fmt"

type Trie struct {
	Children [26]*Trie
	word     string
	mark     bool
}

func (root *Trie) Del(word string) {
	node := root
	for _, c := range word {
		node = node.Children[c-'a']
	}
	node.word = ""
	node.mark = false
}

func findWord(board [][]byte, x, y int, dirs [][]int, root *Trie, visited [][]bool) [][]byte {
	var tmp []byte
	var ans [][]byte
	tmp = append(tmp, board[x][y])
	if root.mark {
		root.mark = false
		root.word = ""
		ans = append(ans, tmp)
	}

	for _, dir := range dirs {
		nx, ny := x+dir[0], y+dir[1]
		if nx >= 0 && nx < len(board) && ny >= 0 && ny < len(board[0]) && root.Children[board[nx][ny]-'a'] != nil && !visited[nx][ny] {
			visited[nx][ny] = true
			nextAnses := findWord(board, nx, ny, dirs, root.Children[board[nx][ny]-'a'], visited)
			visited[nx][ny] = false
			for _, nextAns := range nextAnses {
				tmpSlice := append(nextAns, 0)
				copy(tmpSlice[1:], tmpSlice[0:])
				tmpSlice[0] = board[x][y]
				ans = append(ans, tmpSlice)
			}
		}
	}
	return ans
}

func findWords(board [][]byte, words []string) []string {
	root := &Trie{}

	for _, word := range words {
		node := root
		var acc []byte
		for _, c := range word {
			idx := c - 'a'
			if node.Children[idx] == nil {
				node.Children[idx] = &Trie{}
			}
			node = node.Children[idx]
			acc = append(acc, byte(c))
		}
		node.word = string(acc)
		node.mark = true
	}
	var ans []string
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	duplicate := map[string]bool{}
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if root.Children[board[i][j]-'a'] != nil {
				visited[i][j] = true
				curAnses := findWord(board, i, j, dirs, root.Children[board[i][j]-'a'], visited)
				visited[i][j] = false
				for _, curAns := range curAnses {
					str := string(curAns)
					if !duplicate[str] {
						duplicate[str] = true
						ans = append(ans, str)
					}
				}
			}
		}
	}
	return ans
}

func main() {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	fmt.Println(findWords(board, []string{"oath", "pea", "eat", "rain", "hklf", "hf"}))
}
