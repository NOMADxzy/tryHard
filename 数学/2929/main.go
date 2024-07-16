package main

import "fmt"

func distributeCandies(n int, limit int) int64 {
	ans := int64(0)
	for i := 0; i <= limit; i++ {
		remain23 := n - i
		minVal := max(0, remain23-limit)
		maxVal := min(limit, remain23)
		if minVal <= maxVal {
			ans += int64(maxVal - minVal + 1)
		}
	}
	return ans
}

func main() {
	fmt.Println(distributeCandies(5, 2))
}
