package main

import (
	"fmt"
	"strconv"
)

func add(num1 string, num2 string) string {
	c := 0
	a := num2
	b := num1
	if len(num1) > len(num2) {
		a = num1
		b = num2
	}
	res := ""
	j := len(b) - 1
	var i int
	for i = len(a) - 1; i >= 0 && j >= 0; i-- {
		cur := int(a[i]-'0') + int(b[j]-'0') + c
		j--
		c = cur / 10
		res = strconv.Itoa(cur%10) + res
	}
	for ; i >= 0; i-- {
		cur := int(a[i]-'0') + c
		c = cur / 10
		res = strconv.Itoa(cur%10) + res
	}
	if c > 0 {
		res = "1" + res
	}
	k := 0
	for ; k < len(res)-1 && res[k] == '0'; k++ {
	}
	res = res[k:]
	return res
}

func multiply(num1 string, num2 string) string {
	a := num2
	b := num1
	res := ""
	paddings := 0
	for i := len(a) - 1; i >= 0; i-- {
		cur_s := ""
		c := 0
		for j := len(b) - 1; j >= 0; j-- {
			cur := int(a[i]-'0')*int(b[j]-'0') + c
			c = cur / 10
			cur_s = strconv.Itoa(cur%10) + cur_s
		}
		if c > 0 {
			cur_s = strconv.Itoa(c) + cur_s
		}
		for k := 0; k < paddings; k++ {
			cur_s += "0"
		}
		res = add(res, cur_s)

		paddings++
	}
	return res
}

func main() {
	fmt.Println(multiply("498828660196", "840477629533"))
}
