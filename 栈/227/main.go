package main

import "fmt"

func findNextInt(s string, pos *int) int {
	num := 0
	if s[*pos] == ' ' {
		*pos++
		return findNextInt(s, pos)
	}
	for ; *pos < len(s); *pos++ {
		e := s[*pos]
		if e >= '0' && e <= '9' {
			num = 10*num + int(e-'0')
		} else if e == ' ' {
			continue
		} else {
			break
		}
	}
	return num
}

func calculate(s string) int {
	var stack []int
	pos := 0

	for pos < len(s) {

		//x := findNextInt(s, pos)
		if pos > 0 {
			pos++
			if s[pos-1] == '+' {
				stack = append(stack, findNextInt(s, &pos))
			} else if s[pos-1] == '-' {
				stack = append(stack, -findNextInt(s, &pos))
			} else if s[pos-1] == '*' {
				stack[len(stack)-1] *= findNextInt(s, &pos)
			} else {
				stack[len(stack)-1] /= findNextInt(s, &pos)
			}

		} else {
			stack = append(stack, findNextInt(s, &pos))
		}
	}
	res := 0
	for i := 0; i < len(stack); i++ {
		res += stack[i]
	}
	return res
}

func main() {
	fmt.Println(calculate(" 3+5 / 2 "))
}
