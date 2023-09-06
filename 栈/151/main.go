package main

import "fmt"

func reverseWords(s string) string {
	m := len(s)
	var stack []string

	left := 0
	right := 0

	for right < m {
		if s[left] == ' ' {
			left++
			right++
		} else {
			for right < m && s[right] != ' ' {
				right++
			}
			stack = append(stack, s[left:right])
			right++
			left = right
		}
	}

	resString := ""
	for i := len(stack) - 1; i >= 0; i-- {
		resString += stack[i]
		if i != 0 {
			resString += " "
		}
	}
	return resString
}

func main() {
	fmt.Println(reverseWords("a good   example"))
}
