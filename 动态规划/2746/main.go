package main

import "fmt"

func minimizeConcatenatedLength(words []string) int {
	hist := map[int]int{}
	var dfs func(head, tail byte, pos int) int
	maxPos := len(words) - 1
	maxS := 25

	getHash := func(head, tail byte, pos int) int { // TODO 三维 Hash的计算方法，只有三个变量，系数取上一维度的最大值+1
		s1 := int(head - 'a')
		s2 := int(tail - 'a')
		return (maxS*(maxPos+1)+maxPos+1)*s1 + s2*(maxPos+1) + pos
	}
	dfs = func(head, tail byte, pos int) int {
		if pos == len(words) {
			return 0
		}
		key := getHash(head, tail, pos)
		if v, ok := hist[key]; ok {
			return v
		}
		word := words[pos]
		removeCount1 := dfs(head, word[len(word)-1], pos+1)
		if tail == word[0] {
			removeCount1++
		}
		removeCount2 := dfs(word[0], tail, pos+1)
		if head == word[len(word)-1] {
			removeCount2++
		}
		removeCount := max(removeCount1, removeCount2)
		hist[key] = removeCount
		return removeCount
	}
	sumCount := 0
	for i := 0; i < len(words); i++ {
		sumCount += len(words[i])
	}
	removeCnt := dfs(words[0][0], words[0][len(words[0])-1], 1)
	return sumCount - removeCnt
}

func main() {
	fmt.Println(minimizeConcatenatedLength([]string{"c", "a", "baa"}))
}
