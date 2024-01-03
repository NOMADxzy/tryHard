package main

import "fmt"

func minRemoveToMakeValid(s string) string {
	var stack []int
	mark := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			if len(stack) == 0 {
				mark[i] = true
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	for len(stack) > 0 {
		mark[stack[len(stack)-1]] = true
		stack = stack[:len(stack)-1]
	}
	var ans []byte
	for i := 0; i < len(s); i++ {
		if !mark[i] {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}

func main() {
	fmt.Println(minRemoveToMakeValid("lee(t(c)o)de)"))
}
