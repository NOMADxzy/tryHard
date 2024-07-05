package main

import "fmt"

func wonderfulSubstrings(word string) int64 {
	mask := 1
	valMap := map[byte]int{}
	for i := 0; i < 10; i++ {
		valMap[byte('a'+i)] = mask
		mask *= 2
	}
	cnts := map[int]int64{0: 1}
	cur := 0
	ans := int64(0)
	for i := 0; i < len(word); i++ {
		cur ^= 1 << (word[i] - 'a')
		ans += cnts[cur]
		for j := 0; j < 10; j++ {
			ans += cnts[cur^(1<<j)]
		}
		cnts[cur]++
	}
	return ans
}

func main() {
	fmt.Println(wonderfulSubstrings("aba"))
}
