package main

import "fmt"

func removeDuplicateLetters(s string) string {
	var stack []byte
	m := len(s)
	Cmap := map[byte]int{}

	for i := 0; i < m; i++ {
		Cmap[s[i]] += 1
	}

	for i := 0; i < m; i++ {
		c := s[i]
		if len(stack) == 0 {
			stack = append(stack, c)
		} else {
			exist := false
			for i := 0; i < len(stack); i++ {
				if stack[i] == c {
					exist = true
					break
				}
			}
			if exist {
				Cmap[c] -= 1
				continue
			}
			for len(stack) > 0 && stack[len(stack)-1] > c {
				if Cmap[stack[len(stack)-1]] > 1 {
					Cmap[stack[len(stack)-1]] -= 1
					stack = stack[0 : len(stack)-1]
				} else {
					break
				}
			}
			stack = append(stack, c)
		}
	}
	newS := ""
	for i := 0; i < len(stack); i++ {
		newS += string(stack[i])
	}
	return newS
}

func main() {
	fmt.Println(removeDuplicateLetters("bbcaac"))
}
