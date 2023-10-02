package main

import (
	"fmt"
	"strconv"
)

func findNext(nArr []int, pos int) {
	if pos == len(nArr) {
		return
	}
	if pos == 0 || nArr[pos] >= nArr[pos-1] {
		findNext(nArr, pos+1)
	} else {
		p := pos
		for p > 0 && nArr[p] < nArr[p-1] {
			nArr[p-1]--
			nArr[p] = 9
			p--
		}
		for p = pos + 1; p < len(nArr); p++ {
			nArr[p] = 9
		}
		return
	}
}

func monotoneIncreasingDigits(n int) int {
	wid := len(strconv.Itoa(n))
	nArr := make([]int, wid)
	for i := 0; i < wid; i++ {
		nArr[wid-1-i] = n % 10
		n /= 10
	}
	findNext(nArr, 0)
	sum := 0
	for i := 0; i < wid; i++ {
		sum = 10*sum + nArr[i]
	}
	return sum
}

func main() {
	fmt.Println(monotoneIncreasingDigits(0))
}
