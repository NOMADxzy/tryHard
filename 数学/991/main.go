package main

import "fmt"

func brokenCalc(startValue int, target int) int {
	cur := startValue
	var cnt2, cnt1 int
	mask := 1
	for cur < target {
		cur *= 2
		cnt2++
		mask *= 2
	}
	for cur > target {
		for mask > cur-target {
			mask /= 2
		}
		cur -= mask
		cnt1++
	}
	return cnt1 + cnt2
}

func main() {
	fmt.Println(brokenCalc(2, 3))
}
