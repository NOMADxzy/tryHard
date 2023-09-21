package main

import "fmt"

func countPrimes(n int) int {
	mark := make([]bool, n)
	cnt := 0

	for i := 2; i < n; i++ {
		if mark[i] {
			continue
		} else {
			cnt++
			for k := i * i; k < n; k += i {
				mark[k] = true
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(countPrimes(136649))
}
