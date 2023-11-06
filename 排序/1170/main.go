package main

import (
	"fmt"
	"sort"
)

func f(s string, m map[byte]int) int {
	clear(m)
	curChar := uint8('z')
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		if s[i] < curChar {
			curChar = s[i]
		}
	}
	return m[curChar]
}

func numSmallerByFrequency(queries []string, words []string) []int {
	wordsCnt := make([]int, len(words))
	m := map[byte]int{}
	for i := 0; i < len(words); i++ {
		wordsCnt[i] = f(words[i], m)
	}
	sort.Slice(wordsCnt, func(i, j int) bool {
		return wordsCnt[i] < wordsCnt[j]
	})

	n := len(queries)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		target := f(queries[i], m)
		left, right := 0, len(wordsCnt)-1
		if wordsCnt[left] > target {
			ans[i] = len(wordsCnt)
		} else if wordsCnt[right] <= target {
			ans[i] = 0
		} else {
			for left < right {
				mid := (left + right) / 2
				if wordsCnt[mid] <= target {
					left = mid + 1
				} else {
					right = mid
				}
			}
			ans[i] = len(wordsCnt) - right
		}
	}
	return ans
}

func main() {
	fmt.Println(numSmallerByFrequency([]string{"bbb", "cc"}, []string{"a", "aa", "aaa", "aaaa"}))
}
