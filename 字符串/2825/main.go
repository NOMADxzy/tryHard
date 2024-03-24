package main

import "fmt"

func canMakeSubsequence(str1 string, str2 string) bool {
	for d := 1; d <= 1; d++ {
		var i, j int // 字符串子序列问题往双指针上靠 一定是没问题的
		for j < len(str2) && i < len(str1) {
			a, b := str1[i], byte('a'+(str1[i]-'a'+uint8(d))%26)
			if str2[j] == a || str2[j] == b {
				j++
			}
			i++
		}
		if j == len(str2) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(canMakeSubsequence("abc", "ad"))
}
