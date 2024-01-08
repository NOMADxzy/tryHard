package main

import "fmt"

func findLexSmallestString(s string, a int, b int) string {
	cnt1 := 1
	cnt0 := 1
	for (cnt1*a)%10 != 0 {
		cnt1++
	}
	if b%2 == 1 {
		cnt0 = cnt1
	}
	addVal := func(d0 int, d1 int) string {
		res := make([]byte, len(s))
		for i := 0; i < len(s); i++ {
			newVal := int(s[i] - '0')
			if i%2 == 0 {
				newVal += d0
			} else {
				newVal += d1
			}
			newVal = newVal % 10
			res[i] = byte('0' + newVal)
		}
		return string(res)
	}
	best := s
	for i := 0; i < cnt1; i++ {
		for j := 0; j < cnt0; j++ {
			cur := addVal(j*a, i*a)
			for rightStep := 0; rightStep == 0 || (rightStep*b)%len(s) != 0; rightStep++ {
				mid := rightStep * b % len(s)
				newCur := cur[mid:] + cur[:mid]
				if newCur < best {
					best = newCur
				}
			}
		}
	}
	return best
}

func main() {
	fmt.Println(findLexSmallestString("5525", 9, 2))
}
