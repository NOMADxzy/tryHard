package main

import "fmt"

func kthFactor(n int, k int) int {
	hist := make([]int, k)
	idx := 0
	var i int
	for i = 1; idx < k && i*i < n; i++ {
		if n%i == 0 {
			hist[idx] = i
			idx++
		}
	}
	if idx == k {
		return hist[k-1]
	} else {
		cnt := idx * 2
		if i*i == n {
			hist[idx] = i
			cnt++
		}
		if k > cnt {
			return -1
		}
		return n / hist[cnt-k]
	}
}

func main() {
	fmt.Println(kthFactor(4, 2))
}
