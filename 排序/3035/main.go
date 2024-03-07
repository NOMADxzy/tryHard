package main

import (
	"fmt"
	"sort"
)

func maxPalindromesAfterOperations(words []string) int {
	cnts := map[byte]int{}
	lens := make([]int, len(words))
	for j, word := range words {
		for i := 0; i < len(word); i++ {
			cnts[word[i]]++
		}
		lens[j] = len(word)
	}
	cnt1, cnt2 := 0, 0
	for _, m := range cnts {
		cnt2 += m / 2
		if m%2 == 1 {
			cnt1++
		}
	}

	sort.Slice(lens, func(i, j int) bool {
		return lens[i] < lens[j]
	})

	ans := 0
	for i := 0; i < len(lens); i++ {
		size := lens[i]
		if size%2 == 1 && cnt2*2 < size-1 || size%2 == 0 && cnt2*2 < size {
			break
		}
		ans++
		cnt2 -= size / 2
		if size%2 == 1 {
			cnt1--
		}
	}
	return ans
}

func main() {
	fmt.Println(maxPalindromesAfterOperations([]string{"cd", "ef", "a"}))
}
