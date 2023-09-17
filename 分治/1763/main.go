package main

import "fmt"

func max(a string, b string) string {
	if len(a) >= len(b) {
		return a
	} else {
		return b
	}
}

func findNice(s string, start int, end int) string {
	if start >= end {
		return ""
	}
	marks_upper := 0
	marks_lower := 0
	for i := start; i <= end; i++ {
		if s[i] <= 'Z' {
			mask := 1 << int(s[i]-'A')
			if marks_upper&mask == 0 {
				marks_upper += mask
			}
		} else {
			mask := 1 << int(s[i]-'a')
			if marks_lower&mask == 0 {
				marks_lower += mask
			}
		}
	}
	if marks_upper^marks_lower != 0 {
		for i := start; i <= end; i++ {
			if s[i] <= 'Z' {
				if marks_lower&(1<<int(s[i]-'A')) == 0 {
					return max(findNice(s, start, i-1), findNice(s, i+1, end))
				}
			} else {
				if marks_upper&(1<<int(s[i]-'a')) == 0 {
					return max(findNice(s, start, i-1), findNice(s, i+1, end))
				}
			}
		}
	}
	return s[start : end+1]
}

func longestNiceSubstring(s string) string {
	return findNice(s, 0, len(s)-1)
}

func main() {
	fmt.Println(longestNiceSubstring("dDzeE"))
	//fmt.Println(1038 & (1 << int('B'-'A')))
}
