package main

import (
	"fmt"
	"strconv"
)

func punishmentNumber(n int) int {
	ans := 0
	var dfs func(s string, remain int) bool
	dfs = func(s string, remain int) bool {
		if len(s) == 0 {
			return remain == 0
		}
		for j := 1; j <= len(s); j++ {
			cur, _ := strconv.Atoi(s[:j])
			if cur <= remain && dfs(s[j:], remain-cur) {
				return true
			}
		}
		return false
	}
	for i := 1; i <= n; i++ {
		tmp := strconv.Itoa(i * i)
		if dfs(tmp, i) {
			ans += i * i
		}
	}
	return ans
}

func main() {
	fmt.Println(punishmentNumber(10))
}
