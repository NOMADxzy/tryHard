package main

import "fmt"

func maxProduct(words []string) int {
	features := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		feature := 0
		for j := 0; j < len(words[i]); j++ {
			mask := 1 << int(words[i][j]-'a')
			if feature&mask == 0 {
				feature |= mask
			}
		}
		features[i] = feature
	}
	maxAns := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if features[i]&features[j] == 0 {
				maxAns = max(maxAns, len(words[i])*len(words[j]))
			}
		}
	}
	return maxAns
}

func main() {
	fmt.Println(maxProduct([]string{"abcw", "baz", "foo", "bar", "fxyz", "abcdef"}))
}
