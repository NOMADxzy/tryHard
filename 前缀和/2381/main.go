package main

import "fmt"

func shiftingLetters(s string, shifts [][]int) string {
	m := len(s)
	ans := make([]byte, m)
	//cnts := make([]int, m)
	defers := make([]int, m+1)
	for _, shift := range shifts {
		start, end, t := shift[0], shift[1], shift[2]
		if t == 0 {
			t = 25
		}
		defers[start] += t
		defers[end+1] -= t
		//for i := start; i <= end; i++ {
		//	cnts[i] = (cnts[i] + t) % 26
		//}
	}
	cur := 0
	for i := 0; i < m; i++ {
		cur += defers[i]
		delta := uint8(cur % 26)
		if 'z'-s[i] < delta {
			delta -= 26
		}
		ans[i] = s[i] + delta
	}
	return string(ans)
}

func main() {
	fmt.Println(shiftingLetters("abc", [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}}))
}
