package main

import "fmt"

func partitionLabels(s string) []int {
	Kmap := map[byte]int{}

	for i := 0; i < len(s); i++ {
		Kmap[s[i]]++
	}

	var chips []int
	waitings := 0
	lastIdx := -1

	tmpMap := map[byte]int{}
	for i := 0; i < len(s); i++ {
		tmpMap[s[i]]++
		if tmpMap[s[i]] == 1 {
			waitings++
		}
		if tmpMap[s[i]] == Kmap[s[i]] {
			waitings--
		}
		if waitings == 0 {
			chips = append(chips, i-lastIdx)
			lastIdx = i
		}
	}
	return chips
}

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}
