package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	charMap := map[byte]int{}

	for i := 0; i < len(magazine); i++ {
		charMap[magazine[i]]++
	}
	for i := 0; i < len(ransomNote); i++ {
		charMap[ransomNote[i]]--
		if charMap[ransomNote[i]] < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canConstruct("aa", "ab"))
}
