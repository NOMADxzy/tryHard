package main

import "fmt"

func max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func find(s string, l int, r int, k int) int {
	if r-l+1 < k {
		return 0
	}
	Kmap := map[byte]int{}
	for i := l; i <= r; i++ {
		Kmap[s[i]] += 1
	}

	for i := l; i <= r; i++ {
		if Kmap[s[i]] < k {
			return max(find(s, l, i-1, k), find(s, i+1, r, k))
		}
	}
	return r - l + 1
}

func longestSubstring(s string, k int) int {
	return find(s, 0, len(s)-1, k)
}

func main() {
	fmt.Println(longestSubstring("ababbc", 2))
}
