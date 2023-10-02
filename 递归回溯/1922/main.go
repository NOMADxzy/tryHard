package main

import "fmt"

//常规方法，超时
func countGoodNumbers1(n int64) int {
	if n == 1 {
		return 5
	}
	if n%2 == 0 {
		return (4 * countGoodNumbers(n-1) % 1000000007) % 1000000007
	} else {
		return (5 * countGoodNumbers(n-1) % 1000000007) % 1000000007
	}
}

// 二分优化
func countGoodNumbers(n int64) int {
	if n == 2 {
		return 20
	} else if n == 1 {
		return 5
	} else if n == 3 {
		return 100
	}
	if n%4 == 0 {
		m := countGoodNumbers(n / 2)
		return (m * m) % 1000000007
	} else if n%4 == 1 {
		m := countGoodNumbers((n - 1) / 2)
		return (5 * ((m * m) % 1000000007)) % 1000000007
	} else if n%4 == 2 {
		m := countGoodNumbers((n - 2) / 2)
		return (20 * ((m * m) % 1000000007)) % 1000000007
	} else {
		m := countGoodNumbers((n - 3) / 2)
		return (100 * ((m * m) % 1000000007)) % 1000000007
	}
}

func main() {
	fmt.Println(countGoodNumbers(806166225460393))
}
