package main

import "fmt"

func isValid(s string) bool {
	var stack []byte
	m := len(s)

	for i := 0; i < m; i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if !(s[i] == ')' && top == '(' || s[i] == '}' && top == '{' || s[i] == ']' && top == '[') {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("(()"))
}
