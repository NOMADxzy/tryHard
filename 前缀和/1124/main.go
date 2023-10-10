package main

import "fmt"

func longestWPI(hours []int) int {
	sumArr := make([]int, len(hours))

	if hours[0] > 8 {
		sumArr[0] = 1
	}
	for i := 1; i < len(hours); i++ {
		sumArr[i] = sumArr[i-1]
		if hours[i] > 8 {
			sumArr[i]++
		}
	}

	maxLen := 0
	for i := 0; i < len(hours); i++ {
		sumArr[i] = 2*sumArr[i] - i
		if sumArr[i] > 1 {
			maxLen = i + 1
			continue
		}
		for j := 0; j < i-maxLen; j++ {
			if sumArr[i] > sumArr[j] {
				if i-j > maxLen {
					maxLen = i - j
				} else {
					break
				}
			}
		}
	}
	return maxLen
}

func main() {
	fmt.Println(longestWPI([]int{9, 9, 6, 0, 6, 6, 9}))
}
