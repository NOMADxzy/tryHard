package main

import "fmt"

func maxUniqueSplit(s string) int {
	m := len(s)
	ans := 0
	var dfs func(pos, acc int)
	sMap := map[string]bool{}
	dfs = func(pos, acc int) {
		if pos == m {
			ans = max(ans, acc)
			return
		}
		for right := pos; right < m && acc+m-right > ans; right++ {
			if !sMap[s[pos:right+1]] {
				sMap[s[pos:right+1]] = true
				dfs(right+1, acc+1)
				sMap[s[pos:right+1]] = false
			}
		}
	}
	dfs(0, 0)
	return ans
}

func main() {
	fmt.Println(maxUniqueSplit("wwwzfvedwfvhsww"))
}
