package main

import "fmt"

func findRepeatedDnaSequences1(s string) []string {
	stringMap := map[string]int{}
	var results []string

	for i := 0; i <= len(s)-10; i++ {
		part := s[i : i+10]
		stringMap[part] += 1
		if stringMap[part] == 2 {
			results = append(results, part)
		}
	}
	return results
}

func findRepeatedDnaSequences(s string) []string {
	cnts := map[int]int{}
	s9 := 1
	for i := 0; i < 10; i++ {
		s9 *= 4
	}
	hash := 0
	index := func(b byte) int {
		if b == 'A' {
			return 0
		} else if b == 'T' {
			return 1
		} else if b == 'C' {
			return 2
		} else {
			return 3
		}
	}
	var ans []string
	for i := 0; i < len(s); i++ {
		hash = hash*4 + index(s[i])
		if i > 9 {
			hash -= s9 * index(s[i-10])
		}
		if i >= 9 {
			cnts[hash]++
			if cnts[hash] == 2 {
				ans = append(ans, s[i-9:i+1])
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
}
