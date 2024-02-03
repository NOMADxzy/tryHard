package main

import (
	"fmt"
	"strconv"
)

func simplifiedFractions(n int) []string {
	var ans []string
	mark := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		mark[i] = map[int]bool{}
	}
	for x := 1; x < n; x++ {
		for y := x + 1; y <= n; y++ {
			if !mark[y-1][x] {
				ans = append(ans, strconv.Itoa(x)+"/"+strconv.Itoa(y))
				for j := 2; j*y <= n; j++ {
					mark[j*y-1][j*x] = true
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(simplifiedFractions(4))
}
