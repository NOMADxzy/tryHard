package main

import "fmt"

func countSubstrings(s string) int {
	preAns := []int{0}

	ans := 1

	for i := 1; i < len(s); i++ {
		var curAns []int
		curAns = append(curAns, i)
		if s[i] == s[i-1] {
			curAns = append(curAns, i-1)
		}
		for j := 0; j < len(preAns); j++ {
			pos := preAns[j] - 1
			if pos >= 0 && s[i] == s[pos] {
				curAns = append(curAns, pos)
			}
		}
		ans += len(curAns)
		preAns = curAns
		curAns = nil
	}
	return ans
}

func main() {
	fmt.Println(countSubstrings("aaa"))
}
