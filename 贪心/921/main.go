package main

import "fmt"

func minAddToMakeValid(s string) int {
	var stack int
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack++
		} else if stack > 0 {
			stack--
		} else {
			ans++
		}
	}
	ans += stack
	return ans
}

func main() {
	fmt.Println(minAddToMakeValid("())"))
}
