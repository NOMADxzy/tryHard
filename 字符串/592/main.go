package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getCommonMul(a, b int) int {
	m := a * b
	if a < b {
		a, b = b, a
	}
	div := 0
	for b != 0 {
		remain := a % b
		if remain == 0 {
			div = b
		}
		a = b
		b = remain
	}
	return m / div
}

func fractionAddition(expression string) string {
	var splits [][]int
	expresses := strings.FieldsFunc(expression, func(r rune) bool {
		return r == '+' || r == '-'
	})
	var symbs []bool
	if expression[0] != '-' {
		symbs = append(symbs, true)
	}
	for i := 0; i < len(expression); i++ {
		if expression[i] == '+' {
			symbs = append(symbs, true)
		} else if expression[i] == '-' {
			symbs = append(symbs, false)
		}
	}
	common := 1
	for i, express := range expresses {
		e := strings.Split(express, "/")
		a, _ := strconv.Atoi(e[0])
		b, _ := strconv.Atoi(e[1])
		if !symbs[i] {
			a = -a
		}
		common = getCommonMul(common, b)
		a *= common / b
		splits = append(splits, []int{a, common})
	}
	ans := 0
	for i := 0; i < len(splits); i++ {
		ans += splits[i][0] * (common / splits[i][1])
	}
	if ans == 0 {
		return "0/1"
	}
	var div int
	if ans < 0 {
		div = -ans * common / getCommonMul(-ans, common)
	} else {
		div = ans * common / getCommonMul(ans, common)
	}
	return strconv.Itoa(ans/div) + "/" + strconv.Itoa(common/div)
}

func main() {
	fmt.Println(fractionAddition("-1/2+1/2"))
}
