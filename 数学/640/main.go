package main

import (
	"fmt"
	"strconv"
	"strings"
)

func find(s string) (int, int) {
	xCnt := 0
	rVal := 0
	var symb []bool
	if s[0] != '-' {
		symb = append(symb, true)
	}
	vals := strings.FieldsFunc(s, func(r rune) bool {
		return r == '-' || r == '+'
	})
	for i := 0; i < len(s); i++ {
		if s[i] == '-' {
			symb = append(symb, false)
		} else if s[i] == '+' {
			symb = append(symb, true)
		}
	}
	for i, val := range vals {
		if val[len(val)-1] == 'x' {
			c := 1
			if len(val) > 1 {
				c, _ = strconv.Atoi(val[:len(val)-1])
			}
			if !symb[i] {
				c = -c
			}
			xCnt += c
		} else {
			v, _ := strconv.Atoi(val)
			if !symb[i] {
				v = -v
			}
			rVal += v
		}
	}
	return xCnt, rVal
}

func solveEquation(equation string) string {
	equation_split := strings.Split(equation, "=")
	leftS, RightS := equation_split[0], equation_split[1]
	x1, v1 := find(leftS)
	x2, v2 := find(RightS)
	x := x1 - x2
	v := v2 - v1
	if x == 0 {
		if v == 0 {
			return "Infinite solutions"
		} else {
			return "No solution"
		}
	} else if v%x != 0 {
		return "No solution"
	} else {
		return "x=" + strconv.Itoa(v/x)
	}
}

func main() {
	fmt.Println(solveEquation("-2x+5-3+x=-6+x-2"))
}
