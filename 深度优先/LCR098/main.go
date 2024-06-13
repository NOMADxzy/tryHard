package main

import "fmt"

func uniquePaths(m int, n int) int {
	m--
	n--
	hist := map[int]int{m*100 + n: 1}
	var dfs func(curM, curN int) int
	dfs = func(curM, curN int) int {
		key := curM*100 + curN
		if v, ok := hist[key]; ok {
			return v
		}
		cnt := 0
		if curM < m {
			cnt += dfs(curM+1, curN)
		}
		if curN < n {
			cnt += dfs(curM, curN+1)
		}
		hist[key] = cnt
		return cnt
	}
	return dfs(0, 0)
}

func main() {
	fmt.Println(uniquePaths(3, 7))
}
