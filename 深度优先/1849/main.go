package main

import (
	"fmt"
	"strconv"
)

func splitString(s string) bool {
	var dfs func(target, i int) bool
	dfs = func(target int, pos int) bool {
		if pos == len(s) {
			return true
		}
		for i := pos + 1; i <= len(s); i++ {
			v, _ := strconv.Atoi(s[pos:i])
			if v > target {
				break
			} else if v == target && dfs(target-1, i) {
				return true
			}
		}
		return false
	}
	if len(s) == 1 {
		return false
	}
	for i := 0; i < len(s)-1; i++ {
		v, _ := strconv.Atoi(s[:i+1])
		if dfs(v-1, i+1) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(splitString("050043"))
}
