package main

import "fmt"

func maxLen(s string, k int) int {
	m := len(s)
	left, right := 0, m
	check := func(limit int) bool {
		cnt := 0
		for i := 0; i < m && cnt <= k; i++ {
			for i < m && s[i] == '0' {
				i++
			}
			j := i
			for j < m && s[j] == '1' && j-i+1 <= limit {
				j++
			}
			if j == m {
				break
			} else if s[j] == '0' {
				i = j
				continue
			}
			cnt++
			i = j
		}
		return cnt <= k
	}
	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func main() {
	fmt.Println(maxLen("101110110", 1))
}
