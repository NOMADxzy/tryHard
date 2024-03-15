package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solve(mask, a, b string) int {
	sliceToInts := func(s string) []int {
		ans := make([]int, 4)
		tmp := strings.Split(s, ".")
		if len(tmp) != 4 {
			return []int{}
		}
		for i, str := range tmp {
			val, err := strconv.Atoi(str)
			if err != nil || val > 255 || val < 0 {
				return []int{}
			}
			ans[i] = val
		}
		return ans
	}
	maskArr := sliceToInts(mask)
	aArr := sliceToInts(a)
	bArr := sliceToInts(b)
	aNet, bNet := make([]int, 4), make([]int, 4)
	if len(maskArr) != 4 || len(aArr) != 4 || len(bArr) != 4 {
		return 1
	}

loop:
	for i := 0; i < 4; i++ {
		m := 128
		for m > 0 {
			if maskArr[i]&m == 0 {
				if !(maskArr[0] == 0 && maskArr[1] == 0 && maskArr[2] == 0 && maskArr[3] == 0) {
					return 1
				}
				break loop
			}
			maskArr[i] -= m
			aNet[i] += aArr[i] & m
			bNet[i] += bArr[i] & m
			m /= 2
		}
	}
	for i := 0; i < 4; i++ {
		if aNet[i] != bNet[i] {
			return 2
		}
	}
	return 0
}

func main() {
	var mask, a, b string
	_, _ = fmt.Scan(&mask, &a, &b)
	fmt.Println(solve(mask, a, b))
}
