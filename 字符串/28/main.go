package main

import "fmt"

func strStr(haystack string, needle string) int {
	l := len(haystack)
	r := len(needle)
	if l == 0 && r == 0 {
		return 0
	} else if l == 0 || r == 0 {
		return -1
	}
	k := 0

	_, nextVal := GetNext(needle)

	for i := 0; i < l; {
		if haystack[i] == needle[k] {
			k++
			i++
		} else if k == 0 {
			i++
		} else {
			k = nextVal[k]
		}
		if k == len(needle) {
			return i - k
		}
	}
	return -1
}

func GetNext(str string) ([]int, []int) {
	m := len(str)
	if m == 1 {
		return []int{-1}, []int{-1}
	}
	next := make([]int, m)
	nextVal := make([]int, m)
	i := 2
	k := 0
	nextVal[0] = -1
	nextVal[1] = 0
	next[0] = -1
	next[1] = 0

	for i < m {
		if str[i-1] == str[k] {
			next[i] = k + 1
			if str[i] != str[k+1] {
				nextVal[i] = next[i]
			} else {
				nextVal[i] = nextVal[next[i]]
			}
			k++
			i++
		} else if k == 0 {
			next[i] = 0
			nextVal[i] = 0
			i++
		} else {
			k = nextVal[k]
		}
	}
	return next, nextVal
}

func main() {
	fmt.Println(strStr("a", "a"))
}
