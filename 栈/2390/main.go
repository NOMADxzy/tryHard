package main

import "fmt"

func removeStars(s string) string {
	str := ""

	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			str = str[:len(str)-1]
		} else {
			str += string(s[i])
		}
	}

	return str
}

func main() {
	fmt.Println(removeStars("leet**cod*e"))
}
