package main

import "fmt"

func minSwaps(s string) int {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			stack = append(stack, '[')
		} else {
			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, s[i])
			}
		}
	}
	pairs := len(stack) / 2
	ans := pairs / 2
	if pairs%2 == 1 {
		ans++
	}
	return ans
}

func main() {
	fmt.Println(minSwaps("][]["))
}
