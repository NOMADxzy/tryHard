package main

import "fmt"

func reverse(s string) string {
	arr := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		arr[i] = s[len(arr)-1-i]
	}
	return string(arr)
}

func reverseParentheses(s string) string {
	var stack []string
	var newS string

	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			seq := ""
			for stack[len(stack)-1] != "(" {
				seq += reverse(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			stack = append(stack, seq)
		} else {
			stack = append(stack, s[i:i+1])
		}
	}
	for _, s2 := range stack {
		newS += s2
	}
	return newS
}

func main() {
	fmt.Println(reverseParentheses("an(ed(et(oc))el)"))
}
