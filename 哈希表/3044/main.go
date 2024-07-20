package main

import "fmt"

func mostFrequentPrime(mat [][]int) int {
	primes := map[int]int{-1: 0}
	isPrime := func(x int) bool {
		if primes[x] > 0 {
			return true
		}
		for i := 2; i*i <= x; i++ {
			if x%i == 0 {
				return false
			}
		}
		return true
	}
	ans := -1

	dirs := [][]int{{1, 1}, {1, 0}, {1, -1}, {0, 1}, {0, -1}, {-1, 1}, {-1, 0}, {-1, -1}}
	m, n := len(mat), len(mat[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for _, dir := range dirs {
				p, q := i, j
				cur := mat[p][q]
				for {
					np, nq := p+dir[0], q+dir[1]
					if np >= 0 && np < m && nq >= 0 && nq < n {
						p, q = np, nq
					} else {
						break
					}
					cur = cur*10 + mat[p][q]
					if isPrime(cur) {
						primes[cur]++
						if primes[cur] > primes[ans] || (primes[cur] == primes[ans] && cur > ans) {
							ans = cur
						}
					}
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(mostFrequentPrime([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
