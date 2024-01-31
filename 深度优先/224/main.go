package main

import (
	"fmt"
	"strings"
)

func calculate(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	rightBracketPos := map[int]int{}
	var stack []int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			rightBracketPos[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
	}

	var getVal func(start, end int) int
	getVal = func(start, end int) int {
		sum := 0
		for i := start; i <= end; {
			symb := 1
			if s[i] == '-' {
				symb = -1
				i++
			} else if s[i] == '+' {
				i++
			}
			if s[i] == '(' {
				val := getVal(i+1, rightBracketPos[i]-1)
				val *= symb
				sum += val
				i = rightBracketPos[i] + 1
			} else {
				val := 0
				var j int
				for j = i; j <= end && (s[j] >= '0' && s[j] <= '9'); j++ {
					val = val*10 + int(s[j]-'0')
				}
				val *= symb
				sum += val
				i = j
			}
		}
		return sum
	}
	return getVal(0, len(s)-1)
}

func main() {
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
}
