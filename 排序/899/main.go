package main

import (
	"fmt"
	"sort"
)

func orderlyQueue(s string, k int) string {
	if k == 1 {
		best := s
		for i := 0; i < len(s); i++ {
			if best > s[i:]+s[:i] {
				best = s[i:] + s[:i]
			}
		}
		return best
	} else {
		barr := make([]byte, len(s))
		for i := 0; i < len(s); i++ {
			barr[i] = s[i]
		}
		sort.Slice(barr, func(i, j int) bool {
			return barr[i] < barr[j]
		})
		return string(barr)
	}
}

func main() {
	fmt.Println(orderlyQueue("bac", 1))
}
