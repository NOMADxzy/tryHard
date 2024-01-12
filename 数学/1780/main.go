package main

import "fmt"

func checkPowersOfThree(n int) bool {
	if n%3 != 0 {
		n--
	}
	for n > 0 {
		if n%3 != 0 {
			return false
		} else {
			n /= 3
		}
		if n%3 != 0 {
			n--
		}
	}
	return true
}

func main() {
	fmt.Println(checkPowersOfThree(91))
}
