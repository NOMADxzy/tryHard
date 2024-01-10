package main

import "fmt"

func decode(encoded []int) []int {
	n := len(encoded) + 1
	ans := make([]int, n)
	first := 0
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] ^ encoded[i-1]
		first ^= ans[i]
	}
	if (n-1)/2%2 == 0 { //共有偶数个 偶数 . 偶数+1 对，最终总异或结果是 1 即 （1 异或 trueFirst = first）
		first = first ^ 1
	}
	for i := 0; i < n; i++ {
		ans[i] = ans[i] ^ first
	}
	return ans
}

func main() {
	fmt.Println(decode([]int{6, 5, 4, 6}))
}
