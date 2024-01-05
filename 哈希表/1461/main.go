package main

import "fmt"

func hasAllCodes(s string, k int) bool {
	sMap := map[string]bool{}
	cnt := 0
	target := 1
	for i := 0; i < k; i++ {
		target *= 2
	}
	for i := 0; i <= len(s)-k && cnt < target; i++ {
		if !sMap[s[i:i+k]] {
			sMap[s[i:i+k]] = true
			cnt++
		}
	}
	return cnt == target
}

func main() {
	fmt.Println(hasAllCodes("01100", 2))
}
