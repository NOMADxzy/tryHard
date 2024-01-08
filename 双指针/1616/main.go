package main

import "fmt"

func checkPalindromeFormation(a string, b string) bool {
	n := len(a)
	if n%2 == 1 {
		half := n / 2
		a = a[:half] + a[half+1:]
		b = b[:half] + b[half+1:]
		n--
	}
	half := n / 2
	aMidLeft := half - 1
	for aMidLeft >= 0 && a[aMidLeft] == a[n-1-aMidLeft] {
		aMidLeft--
	}
	bMidLeft := half - 1
	for bMidLeft >= 0 && b[bMidLeft] == b[n-1-bMidLeft] {
		bMidLeft--
	}
	midLeft := min(aMidLeft, bMidLeft)
	var i int
	type1, type2 := true, true
	for i = 0; i <= midLeft && (type1 || type2); i++ {
		if a[i] != b[n-1-i] {
			type1 = false
		}
		if b[i] != a[n-1-i] {
			type2 = false
		}
	}
	return type1 || type2
}

func main() {
	fmt.Println(checkPalindromeFormation("abcdedcba", "jkjfefmba"))
}
