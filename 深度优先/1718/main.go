package main

import "fmt"

func constructDistancedSequence(n int) []int {
	ans := make([]int, 2*n-1)
	var dfs func(pos int) bool
	marks := make([]bool, n+1)
	dfs = func(pos int) bool {
		if pos == len(ans) {
			return true
		} else if ans[pos] > 0 {
			return dfs(pos + 1)
		}
		for i := n; i > 1; i-- {
			if !marks[i] && i+pos < len(ans) && ans[i+pos] == 0 {
				marks[i] = true
				ans[pos] = i
				ans[i+pos] = i
				if dfs(pos + 1) {
					return true
				}
				marks[i] = false
				ans[pos] = 0
				ans[i+pos] = 0
			}
		}
		if !marks[1] {
			marks[1] = true
			ans[pos] = 1
			if dfs(pos + 1) {
				return true
			}
			ans[pos] = 0
			marks[1] = false
		}
		return false
	}
	dfs(0)
	return ans
}

func main() {
	fmt.Println(constructDistancedSequence(20))
}
