package main

import "fmt"

func candy(ratings []int) int {
	m := len(ratings)
	candys := make([]int, m)
	for i := 0; i < m; i++ {
		candys[i] = 1
	}
	for i := 1; i < m; i++ {
		if ratings[i] > ratings[i-1] {
			candys[i] = candys[i-1] + 1
		}
	}
	for i := m - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candys[i] = max(candys[i+1]+1, candys[i])
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		ans += candys[i]
	}
	return ans
}

func main() {
	fmt.Println(candy([]int{1, 1, 0, 0, 2, 3, 3, 1}))
}
