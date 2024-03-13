package main

import "fmt"

func minimumSteps(s string) int64 {
	shouldPlace := len(s) - 1
	ans := int64(0)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '1' {
			ans += int64(shouldPlace - i)
			shouldPlace--
		}
	}
	return ans
}

func main() {
	fmt.Println(minimumSteps("1010"))
}
