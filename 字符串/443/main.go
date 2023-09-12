package main

import "fmt"

func compress(chars []byte) int {
	m := len(chars)
	var stack []byte

	accLen := 0
	flag := m == 1 || chars[m-1] != chars[m-2]

	for i := 1; i < m; i++ {
		l := 1
		for i < m && chars[i] == chars[i-1] {
			i++
			l++
		}
		if l == 1 {
			chars[accLen] = chars[i-1]
			accLen++
		} else {
			chars[accLen] = chars[i-1]
			accLen++

			for l >= 1 {
				stack = append(stack, '0'+byte(l%10))
				l /= 10
			}
			for len(stack) > 0 {
				chars[accLen] = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				accLen++
			}
		}
	}

	if flag {
		chars[accLen] = chars[m-1]
		accLen++
	}

	fmt.Println(chars)

	return accLen
}

func main() {
	fmt.Println(compress([]byte{'a', 'b', 'b', 'c'}))
}
