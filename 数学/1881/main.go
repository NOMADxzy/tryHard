package main

import "fmt"

func maxValue(n string, x int) string {
	neg := n[0] == '-'
	bigger := func(a, b byte) bool {
		if neg {
			return a < b
		} else {
			return a > b
		}
	}
	if neg {
		n = n[1:]
	}
	x_b := byte('0' + x)
	insertIdx := len(n)
	for i := 0; i < len(n); i++ {
		if bigger(x_b, n[i]) {
			insertIdx = i
			break
		}
	}
	symb := ""
	if neg {
		symb = "-"
	}
	return symb + n[:insertIdx] + string(x_b) + n[insertIdx:]
}

func main() {
	fmt.Println(maxValue("985554", 5))
}
