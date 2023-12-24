package main

import "fmt"

func scoreOfParentheses(s string) int {
	var stack []int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, -1)
		} else {
			if stack[len(stack)-1] == -1 {
				stack = stack[:len(stack)-1]
				a := 1
				for len(stack) > 0 && stack[len(stack)-1] > 0 {
					a += stack[len(stack)-1]
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, a)
			} else {
				a := 2 * stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				for len(stack) > 0 && stack[len(stack)-1] > 0 {
					a += stack[len(stack)-1]
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, a)
			}
		}
	}
	return stack[0]
}

func main() {
	fmt.Println(scoreOfParentheses("(())"))
}
