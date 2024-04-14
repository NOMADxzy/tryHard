package main

import "fmt"

func check(s string) string {
	var ans []byte
	ans = append(ans, s[0])
	m := len(s)
	for i := 1; i < m; i++ {
		if s[i] == ans[len(ans)-1] {
			if len(ans) >= 2 {
				if ans[len(ans)-2] == ans[len(ans)-1] {
					continue
				} else if len(ans) >= 3 && ans[len(ans)-3] == ans[len(ans)-2] {
					continue
				}
			}
		}
		ans = append(ans, s[i])
	}
	return string(ans)
}

func main() {
	var strs []string
	var num int
	fmt.Scanf("%d", &num)
	for i := 0; i < num; i++ {
		var s string
		fmt.Scanln(&s)
		strs = append(strs, s)
	}
	for _, str := range strs {
		fmt.Println(check(str))
	}
}
