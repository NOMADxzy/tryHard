package main

import "fmt"

func decodeString(s string) string {
	res := ""
	var stack []string

	for i := 0; i < len(s); i++ {
		if s[i] != ']' {
			stack = append(stack, s[i:i+1])
		} else {
			tmp := ""
			for len(stack) > 0 && stack[len(stack)-1] != "[" {
				tmp = stack[len(stack)-1] + tmp
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				stack = append(stack, tmp)
			} else {
				stack = stack[:len(stack)-1]
				times := 0
				q := 1
				for len(stack) > 0 && stack[len(stack)-1] <= "9" {
					times += int(stack[len(stack)-1][0]-'0') * q
					stack = stack[:len(stack)-1]
					q *= 10
				}
				newTmp := ""
				for i := 0; i < times; i++ {
					newTmp += tmp
				}
				stack = append(stack, newTmp)
			}

		}
	}
	for i := 0; i < len(stack); i++ {
		res += stack[i]
	}
	return res

}

func main() {
	fmt.Println(decodeString("2[abc]3[cd]ef"))
	//fmt.Println("a" > "1")
}
