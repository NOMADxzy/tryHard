package main

import "fmt"

func smallestNumber(pattern string) string {
	m := len(pattern)
	ans := make([]byte, m+1)
	var stack []int
	idx := 0
	for i := 0; i < m; i++ {
		stack = append(stack, i+1)
		if pattern[i] == 'I' {
			for len(stack) > 0 {
				e := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				ans[idx] = byte('0' + e)
				idx++
			}
		} else {
			continue
		}
	}
	stack = append(stack, m+1)
	for len(stack) > 0 {
		e := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans[idx] = byte('0' + e)
		idx++
	}
	return string(ans)
}

func main() {
	fmt.Println(smallestNumber("IIIDIDDD"))
}
