package main

import "fmt"

func clumsy(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else if n == 3 {
		return 6
	}
	ans := 0
	for i := n; i-3 >= n%4; i -= 4 {
		val := i * (i - 1) / (i - 2)
		if i < n {
			val = -val
		}
		ans += val
	}
	for i := n - 3; i >= n%4 && i > 0; i -= 4 {
		ans += i
	}
	i := n % 4
	val := i
	if i > 2 {
		val *= i - 1
	}
	return ans - val
}

func main() {
	fmt.Println(clumsy(7))
}
