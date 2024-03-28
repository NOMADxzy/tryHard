package main

import "fmt"

// 子串的定义：从字符串中拿出连续的一段，即为子串。例如 bc 是 abc 的子串，但是 ac 不是，因为 ac 在 abc 中不连续。
//
//回文串的定义：正着读和反着读一样的字符串为回文串。例如 noon，bb 等，单个字符也是回文串。
//
//回文子串的定义：是原字符串的一个子串，且是回文串。
//
//给出一个字符串，你可以任意排列里面的字符，使得其回文子串的数量尽可能的多，求回文子串数量最多为多少。

func solve(s string) int {
	cnts := map[byte]int{}
	for i := 0; i < len(s); i++ {
		cnts[s[i]]++
	}
	ans := len(s)
	for _, cnt := range cnts {
		ans += cnt - 1
	}
	return ans
}

func main() {
	var s string
	_, _ = fmt.Scan(&s)
	fmt.Println(solve(s))
}
