package main

import (
	"fmt"
	"sort"
)

func findLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		return len(dictionary[i]) > len(dictionary[j]) || len(dictionary[i]) == len(dictionary[j]) && dictionary[i] < dictionary[j]
	})
	for _, t := range dictionary {
		i, j := 0, 0
		for i < len(s) && j < len(t) {
			if s[i] == t[j] {
				i++
				j++
			} else {
				i++
			}
		}
		if j == len(t) {
			return t
		}
	}
	return ""
}

func main() {
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
}
