package main

import "fmt"

func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}
	plus := 0
	mask := 1
	ans := make([]byte, 50)
	for i := 0; mask <= n; i++ {
		if n&mask > 0 {
			if i%2 == 1 {
				plus += mask
				n += mask
			}
			ans[i] = '1'
		} else {
			ans[i] = '0'
		}
		mask *= 2
	}
	s := ""
	for ans[len(ans)-1] == 0 {
		ans = ans[:len(ans)-1]
	}
	for i := len(ans) - 1; i >= 0; i-- {
		s += string(ans[i])
	}
	return s
}

func main() {
	fmt.Println(baseNeg2(0))
}
