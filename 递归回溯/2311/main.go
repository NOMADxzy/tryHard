package main

import (
	"fmt"
)

func longestSubsequence(s string, k int) int {
	count := 0
	var arr1 []int

	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			count++
		} else {
			arr1 = append(arr1, len(s)-i-1)
		}
	}

	sum := 0
	for i := len(arr1) - 1; i >= 0; i-- {
		if arr1[i] >= 31 {
			break
		}
		sum += 1 << arr1[i]
		if sum <= k {
			count++
		} else {
			break
		}
	}
	return count

}

func main() {
	fmt.Println(longestSubsequence("001010101011010100010101101010010", 93951055))
}
