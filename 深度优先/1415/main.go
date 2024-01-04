package main

import "fmt"

func getHappyString(n int, k int) string {
	rank := 0
	chars := []byte{'a', 'b', 'c'}
	ans := ""
	var dfs func(acc string)
	dfs = func(acc string) {
		if len(ans) > 0 {
			return
		}
		if len(acc) == n {
			rank++
			if rank == k {
				ans = acc
			}
			return
		}
		for _, c := range chars {
			if len(acc) == 0 || c != acc[len(acc)-1] {
				dfs(acc + string(c))
			}
		}
	}
	dfs("")
	return ans
}

func main() {
	fmt.Println(getHappyString(10, 100))
}
