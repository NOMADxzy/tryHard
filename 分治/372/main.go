package main

import "fmt"

func Pow(a int, b int) int {
	r := 1
	for ; b > 0; b /= 2 {
		if b&1 > 0 {
			r = r * a % 1337
		}
		a = a * a % 1337
	}
	return r
}

func superPow(a int, b []int) int {
	sum := 1

	for i := len(b) - 1; i >= 0; i-- {
		sum = (sum * (Pow(a, b[i]) % 1337)) % 1337
		a = Pow(a, 10) % 1337
	}
	return sum
}

func main() {
	fmt.Println(superPow(2147483647, []int{2, 0, 0}))
}
