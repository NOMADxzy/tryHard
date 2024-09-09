package main

import "fmt"

func solve(s string) bool {
	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			cnt++
		} else if cnt == 0 {
			return false
		} else {
			cnt--
		}
	}
	return cnt == 0
}

func main() {
	//fmt.Println(solve("abaababb"))
	var n int
	var s string
	_, _ = fmt.Scan(&n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&s)
		ans := solve(s)
		if ans {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}
