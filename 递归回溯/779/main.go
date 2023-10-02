package main

import "fmt"

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}
	if k%2 == 0 {
		return 1 - kthGrammar(n-1, (k+1)/2)
	} else {
		return kthGrammar(n-1, (k+1)/2)
	}
}

func main() {
	fmt.Println(kthGrammar(2, 2))
}
