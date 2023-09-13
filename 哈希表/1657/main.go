package main

import "fmt"

func closeStrings(word1 string, word2 string) bool {
	map1 := map[byte]int{}
	map2 := map[byte]int{}
	Kmap1 := map[int]int{}
	Kmap2 := map[int]int{}
	for i := 0; i < len(word1); i++ {
		map1[word1[i]] += 1
	}

	for i := 0; i < len(word2); i++ {
		map2[word2[i]] += 1
		if map1[word2[i]] == 0 {
			return false
		}
	}

	for _, num := range map1 {
		Kmap1[num]++
	}

	for _, num := range map2 {
		Kmap2[num]++
	}

	for key, time := range Kmap1 {
		if time != Kmap2[key] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(closeStrings("aaaaabbcccdddeeeeffff", "aaabbbbccddeeeeefffff"))
}
