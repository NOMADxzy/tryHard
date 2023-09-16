package main

import (
	"fmt"
	"strconv"
)

func getHint(secret string, guess string) string {
	m := len(guess)
	n := len(secret)

	length := m
	if n < m {
		length = n
	}

	A, B := 0, 0
	Smap := map[byte]int{}
	Gmap := map[byte]int{}
	var i int
	for i = 0; i < length; i++ {
		if secret[i] == guess[i] {
			A++
		} else {
			Smap[secret[i]]++
			Gmap[guess[i]]++
		}
	}
	for ; i < m; i++ {
		Gmap[guess[i]]++
	}
	for ; i < n; i++ {
		Smap[secret[i]]++
	}

	for b, val := range Smap {
		if Gmap[b] < val {
			val = Gmap[b]
		}
		B += val
	}
	return strconv.Itoa(A) + "A" + strconv.Itoa(B) + "B"
}

func main() {
	fmt.Println(getHint("1", "2"))
}
