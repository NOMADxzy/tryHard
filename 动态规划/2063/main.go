package main

import "fmt"

func countVowels(word string) int64 {
	var ans int64
	preCnts := int64(0)
	for i := 0; i < len(word); i++ {
		newCnts := preCnts
		if word[i] == 'a' || word[i] == 'e' || word[i] == 'i' || word[i] == 'o' || word[i] == 'u' {
			newCnts += int64(i + 1)
		}
		ans += newCnts
		preCnts = newCnts
	}
	return ans
}

func main() {
	fmt.Println(countVowels("aba"))
}
