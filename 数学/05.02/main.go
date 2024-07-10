package main

import "fmt"

func printBin(num float64) string {
	if num == 1 {
		return "1"
	}
	pre := "0."
	vals := ""
	for num > 0 && len(vals) < 32 {
		num *= 2
		if num >= 1 {
			vals += "1"
			num -= 1
		} else {
			vals += "0"
		}
	}
	if num > 0 {
		return "ERROR"
	}
	return pre + vals
}

func main() {
	fmt.Println(printBin(0.625))
}
