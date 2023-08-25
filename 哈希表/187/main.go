package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
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

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAAAA"))
}
