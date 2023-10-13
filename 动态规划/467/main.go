package main

import "fmt"

func findSubstringInWraproundString(s string) int {
	a, b := 1, 0
	ans := 0
	eachMaxLen := make([]int, 26)
	eachMaxLen[int(s[0]-'a')] = 1

	for i := 1; i < len(s); i++ {
		if s[i]-s[i-1] == byte(1) || s[i-1]-s[i] == byte(25) {
			b = a + 1
		} else {
			b = 1
		}
		idx := int(s[i] - 'a')
		if b > eachMaxLen[idx] {
			eachMaxLen[idx] = b
		}
		a = b
	}
	for i := 0; i < 26; i++ {
		ans += eachMaxLen[i]
	}
	return ans
}

func main() {
	fmt.Println(findSubstringInWraproundString("zabce"))
}
