package main

import "fmt"

func goNext(s string, pos int, stack []byte, LeftCnt, RightCnt int) bool {
	if pos == len(s) {
		return len(stack) == 0
	}
	if s[pos] == '(' {
		stack = append(stack, '(')
		return goNext(s, pos+1, stack, LeftCnt+1, RightCnt+1)
	} else if s[pos] == '*' {
		var chars []byte

		if LeftCnt == 0 {
			chars = []byte{' '}
			if RightCnt > 0 {
				chars = []byte{'(', ' '}
			}
		} else if LeftCnt < RightCnt {
			chars = []byte{'(', ' ', ')'}
		} else if LeftCnt == RightCnt {
			chars = []byte{' ', '(', ')'}
		} else if LeftCnt > RightCnt {
			chars = []byte{')', ' ', '('}
		}

		for _, c := range chars {
			if c == ')' && len(stack) > 0 && goNext(s, pos+1, stack[:len(stack)-1], LeftCnt-1, RightCnt) {
				return true
			}
			if c == ' ' && goNext(s, pos+1, stack, LeftCnt, RightCnt) {
				return true
			}
			if c == '(' && goNext(s, pos+1, append(stack, '('), LeftCnt+1, RightCnt) {
				return true
			}
		}

	} else {
		if len(stack) > 0 {
			return goNext(s, pos+1, stack[:len(stack)-1], LeftCnt-1, RightCnt-1)
		}
	}
	return false
}

func checkValidString(s string) bool {
	var stack []byte
	LeftCnt, RightCnt := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			LeftCnt++
		} else if s[i] == ')' {
			RightCnt++
		}
	}
	return goNext(s, 0, stack, 0, RightCnt-LeftCnt)
}

func main() {
	fmt.Println(checkValidString("**************************************************))))))))))))))))))))))))))))))))))))))))))))))))))"))
}
