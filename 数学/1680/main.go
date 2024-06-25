package main

import "fmt"

func concatenatedBinary(n int) int {
	MOD := 1000000007
	acc := 0
	w := 2
	for i := 1; i <= n; i++ {
		if i&w > 0 {
			w *= 2
		}
		acc = (acc*w + i) % MOD
	}
	return acc
}

func main() {
	fmt.Println(concatenatedBinary(3))
}
