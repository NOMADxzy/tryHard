package main

import (
	"fmt"
	"strconv"
)

func cur1(x int) int {
	if x == 1 {
		return 1
	} else {
		return 0
	}
}

func findNext(num string, pos int, preLimit bool, pre1sums int, hist map[int]int) int {
	keyNum := pos*10 + pre1sums
	if preLimit {
		keyNum *= -1
	}
	if ans, ok := hist[keyNum]; ok {
		return ans
	}
	if pos == len(num) {
		return pre1sums
	}
	cnts := 0
	for i := 0; i <= 9; i++ {
		if preLimit && i > int(num[pos]-'0') {
			break
		}
		cnts += findNext(num, pos+1, preLimit && i == int(num[pos]-'0'), pre1sums+cur1(i), hist)
	}
	hist[keyNum] = cnts
	return cnts
}

func countDigitOne(n int) int {
	num := strconv.Itoa(n)
	ans := findNext(num, 0, true, 0, map[int]int{})
	return ans
}

func main() {
	fmt.Println(countDigitOne(824883294))
}
