package main

import "fmt"

func generateParenthesis(n int) []string {
	n *= 2
	var ans []string
	var dfs func(leftCnt, pos int, acc string)
	dfs = func(leftCnt, pos int, acc string) {
		if pos == n {
			ans = append(ans, acc)
			return
		}
		if leftCnt == 0 {
			dfs(leftCnt+1, pos+1, acc+"(")
		} else {
			dfs(leftCnt-1, pos+1, acc+")")
			if n-pos-1 >= leftCnt+1 {
				dfs(leftCnt+1, pos+1, acc+"(")
			}
		}
	}
	dfs(0, 0, "")
	return ans
}

func main() {
	fmt.Println(generateParenthesis(3))
}
