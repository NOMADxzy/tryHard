package main

import "fmt"

func equalSubstring(s string, t string, maxCost int) int {
	length := len(s)
	if len(t) < length {
		length = len(t)
	}
	costs := make([]int, length)
	for i := 0; i < length; i++ {
		if s[i] > t[i] {
			costs[i] = int(s[i] - t[i])
		} else {
			costs[i] = int(t[i] - s[i])
		}
	}
	l, r := 0, 0
	maxLen := 0
	sum := 0
	for ; r < length; r++ {
		sum += costs[r]
		for sum > maxCost && r < length {
			sum -= costs[l]
			l++
			r++
			if r < length {
				sum += costs[r]
			}
		}
		if r == length {
			r--
		}
		if r-l+1 > maxLen {
			maxLen = r - l + 1
		}
	}
	return maxLen
}

func main() {
	fmt.Println(equalSubstring("abcd	", "acde", 0))
}
