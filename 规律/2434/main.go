package main

import "fmt"

func robotWithString(s string) string {
	cnts := make([]int, 26)
	for i := 0; i < len(s); i++ {
		idx := int(s[i] - 'a')
		cnts[idx]++
	}
	var stack []byte
	ans := make([]byte, len(s))
	ansIdx := 0
	hasSmaller := func(i int) bool {
		for j := 0; j < i; j++ {
			if cnts[j] > 0 {
				return true
			}
		}
		return false
	}
	for i := 0; i < len(s); i++ {
		idx := int(s[i] - 'a')
		stack = append(stack, s[i])
		cnts[idx]--

		for !hasSmaller(idx) {
			ans[ansIdx] = stack[len(stack)-1]
			ansIdx++
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			idx = int(stack[len(stack)-1] - 'a')
		}
	}
	return string(ans)
}

func main() {
	fmt.Println(robotWithString("zza"))
}
