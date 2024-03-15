package main

import "fmt"

func divisibilityArray(word string, m int) []int {
	pre := 0
	n := len(word)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		val := int(word[i] - '0')
		cur := ((pre*(10%m))%m + val%m) % m
		if cur == 0 {
			ans[i] = 1
		}
		pre = cur
	}
	return ans
}

func main() {
	fmt.Println(divisibilityArray("998244353", 3))
}
