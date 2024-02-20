package main

import "fmt"

func sortVowels(s string) string {
	str5s := make([]int, 10)
	indexOf := func(a byte) int {
		if a == 'A' {
			return 0
		} else if a == 'E' {
			return 1
		} else if a == 'I' {
			return 2
		} else if a == 'O' {
			return 3
		} else if a == 'U' {
			return 4
		} else if a == 'a' {
			return 5
		} else if a == 'e' {
			return 6
		} else if a == 'i' {
			return 7
		} else if a == 'o' {
			return 8
		} else if a == 'u' {
			return 9
		}
		return -1
	}
	for i := 0; i < len(s); i++ {
		if id := indexOf(s[i]); id >= 0 {
			str5s[id]++
		}
	}
	ans := make([]byte, len(s))
	idx := 0
	strs := []byte{'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u'}
	for i := 0; i < 10 && idx < len(s); {
		for idx < len(s) && indexOf(s[idx]) < 0 {
			ans[idx] = s[idx]
			idx++
		}
		if str5s[i] > 0 {
			str5s[i]--
			ans[idx] = strs[i]
			idx++
		}
		if str5s[i] == 0 {
			i++
		}
	}
	for idx < len(s) {
		ans[idx] = s[idx]
		idx++
	}
	return string(ans)
}

func main() {
	fmt.Println(sortVowels("uZcPmqAd"))
}
