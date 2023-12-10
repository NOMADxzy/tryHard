package main

import "fmt"

func findNext(lbCnt, rbCnt int, s string) []string {
	if lbCnt < rbCnt {
		return nil
	} else if s == "" {
		if lbCnt == rbCnt {
			return []string{""}
		} else {
			return nil
		}
	}
	// 确保lbCnt>=rbCnt && len(s)>0
	var ans []string
	if s[0] == '(' {
		nextAnses1 := findNext(lbCnt+1, rbCnt, s[1:])
		for _, nextAns := range nextAnses1 {
			ans = append(ans, s[:1]+nextAns)
		}
		nextAnses2 := findNext(lbCnt, rbCnt, s[1:])
		for _, nextAns := range nextAnses2 {
			ans = append(ans, nextAns)
		}
	} else if s[0] == ')' {
		nextAnses1 := findNext(lbCnt, rbCnt+1, s[1:])
		for _, nextAns := range nextAnses1 {
			ans = append(ans, s[:1]+nextAns)
		}
		nextAnses2 := findNext(lbCnt, rbCnt, s[1:])
		for _, nextAns := range nextAnses2 {
			ans = append(ans, nextAns)
		}
	} else {
		nextAnses := findNext(lbCnt, rbCnt, s[1:])
		for _, nextAns := range nextAnses {
			ans = append(ans, s[:1]+nextAns)
		}
	}
	return ans
}

func removeInvalidParentheses(s string) []string {
	ans0 := findNext(0, 0, s)
	largest := 0
	for _, a := range ans0 {
		if len(a) > largest {
			largest = len(a)
		}
	}
	set := map[string]bool{}
	var ans []string
	for _, a := range ans0 {
		if len(a) == largest && !set[a] {
			ans = append(ans, a)
		}
		set[a] = true
	}
	return ans
}

func main() {
	fmt.Println(removeInvalidParentheses("()())()"))
}
