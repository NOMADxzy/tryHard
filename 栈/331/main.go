package main

import (
	"fmt"
	"strings"
)

func isValidSerialization(preorder string) bool {

	var stack []int
	strs := strings.Split(preorder, ",")
	for i := 0; i < len(strs); i++ {
		if strs[i] != "#" {
			stack = append(stack, 2)
		} else {
			for len(stack) > 0 && stack[len(stack)-1] == 0 {
				stack = stack[:len(stack)-1]
				if len(stack) == 0 || stack[len(stack)-1] == 0 {
					return false
				}
				stack = stack[:len(stack)-1]
			}
			if i != len(strs)-1 {
				stack = append(stack, 0)
			}
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"))
}
