package main

import "fmt"

// 0=<i<j<n的情况
func makeStringsEqual1(s string, target string) bool {
	// 把有后置1的0变成1，把有前置1的1变成0
	if len(s) != len(target) {
		return false
	}
	left1Pos, right1Pos := -1, -1
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			right1Pos = i
			if left1Pos == -1 {
				left1Pos = i
			}
		}
	}
	for i := 0; i < len(s); i++ {
		if s[i] != target[i] { // 需要改变
			if s[i] == '1' && left1Pos == i { // 是最前的1
				return false
			} else if s[i] == '0' && right1Pos < i { // 是没有后置1的0
				return false
			}
		}
	}
	return true
}

func makeStringsEqual(s string, target string) bool {
	sCnt1 := 0
	targetCnt1 := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			sCnt1++
		}
	}
	for i := 0; i < len(target); i++ {
		if target[i] == '1' {
			targetCnt1++
		}
	}
	return targetCnt1 == 0 && sCnt1 == 0 || targetCnt1 > 0 && sCnt1 > 0
}

//0110
//1110
//0010

//0 0 -> 0 0
//0 1 -> 1 1
//1 1 -> 1 0

func main() {
	fmt.Println(makeStringsEqual("0000", "1111"))
}
