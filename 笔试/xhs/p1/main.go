package main

import "fmt"

func solve(ids []string) []string {
	mark := map[string]bool{}

	var ans []string
	for _, id := range ids {
		if !mark[id] {
			ans = append(ans, id)
			mark[id] = true
		}
	}
	return ans
}

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	ids := make([]string, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&ids[i])
	}
	ans := solve(ids)
	for _, an := range ans {
		fmt.Println(an)
	}
}
