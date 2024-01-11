package main

import "fmt"

func countHomogenous(s string) int {
	ans := 0
	MOD := 1000000007
	for i := 0; i < len(s); {
		start := i
		for i < len(s) && s[i] == s[start] {
			i++
		}
		addVal := (i - start + 1) * (i - start) / 2
		ans = (ans + addVal) % MOD
	}
	return ans
}

func main() {
	fmt.Println(countHomogenous("abbcccaa"))
}
