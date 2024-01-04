package main

import "fmt"

func numSteps(s string) int {
	ans := 0
	c := false
	for i := len(s) - 1; i > 0; i-- {
		if c {
			if s[i] == '1' {
				ans += 1
				c = true
			} else {
				ans += 2
				c = true
			}
		} else {
			if s[i] == '1' {
				ans += 2
				c = true
			} else {
				ans += 1
				c = false
			}
		}
	}
	if c {
		ans += 1
	}
	return ans
}

func main() {
	fmt.Println(numSteps("1101"))
}
