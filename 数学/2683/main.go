package main

import "fmt"

func doesValidArrayExist(derived []int) bool {
	m := len(derived)
	constructs := make([]int, m)
	constructs[0] = 0
	cpt := func(a, r int) int {
		if r == 1 {
			return 1 - a
		} else {
			return a
		}
	}
	for i := 1; i < m; i++ {
		r := derived[i-1]
		constructs[i] = cpt(constructs[i-1], r)
	}
	return derived[m-1] == constructs[m-1]^constructs[0]
}

func main() {
	fmt.Println(doesValidArrayExist([]int{1, 1}))
}
