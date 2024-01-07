package main

import "fmt"

func minInsertions(s string) int {
	var stack int
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack++
		} else if i == len(s)-1 || s[i+1] != ')' {
			ans++
			if stack == 0 {
				ans++
				stack++
			}
			stack--
		} else {
			i++
			if stack == 0 {
				ans++
				stack++
			}
			stack--
		}
	}
	ans += stack * 2
	return ans
}

func main() {
	fmt.Println(minInsertions("))())("))
}
