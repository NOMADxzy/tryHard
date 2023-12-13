package main

import "fmt"

func convert(x int) uint16 {
	if x >= 0 {
		return uint16(x)
	} else {
		return uint16(x)
	}
}

func getSum(a int, b int) int {
	A := convert(a)
	B := convert(b)
	mask := uint16(1)
	c := false
	ans := uint16(0)
	for mask != 0 {
		v1, v2 := A&mask, B&mask
		postCnt := 0
		if c {
			postCnt++
		}
		if v1 > 0 {
			postCnt++
		}
		if v2 > 0 {
			postCnt++
		}
		if postCnt%2 != 0 {
			ans += mask
		}
		if postCnt >= 2 {
			c = true
		} else {
			c = false
		}
		mask *= 2
	}

	if ans&(uint16(1)<<15) > 0 {
		mask = uint16(1)
		for mask <= uint16(1)<<14 {
			if ans&mask > 0 {
				ans -= mask
			} else {
				ans += mask
			}
			mask *= 2
		}
		ans++
		ans -= mask
		return -int(ans)
	} else {
		return int(ans)
	}
}

func main() {
	fmt.Println(getSum(990, -991))
}
