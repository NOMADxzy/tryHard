package main

import "fmt"

func minOperations(k int) int {
	var cnt1 int
	ans := k
	for cnt1 <= k/2+1 {
		val := cnt1 + 1
		cnt2 := k/val - 1
		if k%val != 0 {
			cnt2++
		}
		ans = min(ans, cnt1+cnt2)
		cnt1++
	}
	return ans
}

func main() {
	fmt.Println(minOperations(11))
}
