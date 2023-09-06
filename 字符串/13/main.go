package main

import "fmt"

func romanToInt(s string) int {
	sum := 0
	m := len(s)

	valMap := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	for i := 0; i < m; i++ {
		c := s[i]
		if c == 'I' || c == 'X' || c == 'C' {
			if i+1 < m {
				if c == 'I' && s[i+1] == 'V' || c == 'X' && s[i+1] == 'L' || c == 'C' && s[i+1] == 'D' {
					sum += 4 * valMap[c]
					i++
				} else if c == 'I' && s[i+1] == 'X' || c == 'X' && s[i+1] == 'C' || c == 'C' && s[i+1] == 'M' {
					sum += 9 * valMap[c]
					i++
				} else {
					sum += valMap[c]
				}
			} else {
				sum += valMap[c]
			}
		} else {
			sum += valMap[c]
		}
	}
	return sum
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
