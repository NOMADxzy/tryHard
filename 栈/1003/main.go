package main

import "fmt"

func isValid(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] != 'c' {
			stack = append(stack, s[i])
		} else {
			if len(stack) < 2 || stack[len(stack)-1] != 'b' || stack[len(stack)-2] != 'a' {
				return false
			} else {
				stack = stack[:len(stack)-2]
			}
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("aabcbc"))
}
