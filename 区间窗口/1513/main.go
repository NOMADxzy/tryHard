package main

import "fmt"

func numSub(s string) int {
	MOD := 1000000007
	var i int
	ans := 0
	for i < len(s) {
		for i < len(s) && s[i] == '0' {
			i++
		}
		if i == len(s) {
			break
		}
		l := i
		for i < len(s) && s[i] == '1' {
			i++
		}
		n := i - l
		ans = (ans + (n+1)*n/2) % MOD
	}
	return ans
}

func main() {
	fmt.Println(numSub("0110111"))
}
