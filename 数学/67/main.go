package main

import "fmt"

func yihuo(a bool, b bool) bool {
	if a {
		return !b
	} else {
		return b
	}
}

func addBinary(a string, b string) string {
	m, n := len(a), len(b)

	i, j := m-1, n-1
	s := ""
	c := false

	for i >= 0 && j >= 0 {
		x := a[i] == '1'
		y := b[j] == '1'
		i--
		j--

		if yihuo(yihuo(x, y), c) {
			s = "1" + s
		} else {
			s = "0" + s
		}
		if x && y {
			c = true
		} else if !x && !y {
			c = false
		}
	}
	var remain string
	if i >= 0 {
		remain = a[:i+1]
	} else {
		remain = b[:j+1]
	}
	for k := len(remain) - 1; k >= 0; k-- {
		x := remain[k] == '1'
		if yihuo(x, c) {
			s = "1" + s
		} else {
			s = "0" + s
		}
		if x && c {
			c = true
		} else {
			c = false
		}
	}
	if c {
		s = "1" + s
	}
	return s
}

func main() {
	fmt.Println(addBinary("1", "111"))
}
