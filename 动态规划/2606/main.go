package main

import "fmt"

func maximumCostSubstring(s string, chars string, vals []int) int {
	values := make([]int, 26)
	for i := 0; i < 26; i++ {
		values[i] = i + 1
	}
	for i := 0; i < len(chars); i++ {
		idx := int(chars[i] - 'a')
		values[idx] = vals[i]
	}
	var ans int
	preMax := 0
	for i := 0; i < len(s); i++ {
		curIdx := int(s[i] - 'a')
		curMax := values[curIdx]
		if preMax > 0 {
			curMax += preMax
		}
		ans = max(ans, curMax)
		preMax = curMax
	}
	return ans
}

func main() {
	fmt.Println(maximumCostSubstring("adaa", "d", []int{-1000}))
}
