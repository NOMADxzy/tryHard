package main

import "fmt"

func expressiveWords(s string, words []string) int {
	ans := 0
	for _, word := range words {
		i, j := 0, 0
		for i < len(s) && j < len(word) {
			if word[j] != s[i] {
				break
			} else if j+1 == len(word) || word[j] != word[j+1] { //word[j]是单个字符
				if i+1 < len(s) && s[i+1] == s[i] { //s[i]是2个以上连续
					if i+2 >= len(s) || s[i+2] != s[i] { //非三个连续
						break
					} else {
						for i+1 < len(s) && s[i+1] == s[i] {
							i++
						}
					}
				}
			} else { //word[j]是连续字符
				iLeft := i
				jLeft := j
				for i+1 < len(s) && s[i+1] == s[i] {
					i++
				}
				for j+1 < len(word) && word[j+1] == word[j] {
					j++
				}
				if i-iLeft < j-jLeft {
					break
				}
			}
			i++
			j++
		}
		if i == len(s) && j == len(word) {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(expressiveWords("zzzzzyyyyy", []string{"zzyy", "zy", "zyy"}))
}
