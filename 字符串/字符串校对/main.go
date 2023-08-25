package main

import "fmt"

func check(s string) string {
	for i := 0; i < len(s); {
		//s1 += s[i:i+1]
		if i >= 2 && s[i-2] == s[i-1] {
			if s[i] == s[i-1] {
				s = s[:i] + s[i+1:]
			} else if i+1 < len(s) && s[i] == s[i+1] {
				s = s[:i] + s[i+1:]
			} else {
				i++
			}
		} else {
			i++
		}
	}
	return s
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
