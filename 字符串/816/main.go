package main

import "fmt"

func insertPoints(s string) []string {
	var ans []string
	if len(s) == 1 || len(s) > 1 && s[0] != '0' {
		ans = append(ans, s)
	}
	if len(s) == 1 {
		return ans
	}
	if s[0] == '0' {
		if s[len(s)-1] != '0' {
			return append(ans, "0."+s[1:])
		}
	} else if s[len(s)-1] != '0' {
		for l := 1; l < len(s); l++ {
			ans = append(ans, s[:l]+"."+s[l:])
		}
	}
	return ans
}

// (123)
func ambiguousCoordinates(s string) []string {
	var ans []string
	num := s[1 : len(s)-1]
	for i := 1; i < len(num); i++ {
		left, right := num[:i], num[i:]
		lefts := insertPoints(left)
		rights := insertPoints(right)
		for _, s1 := range lefts {
			for _, s2 := range rights {
				ans = append(ans, "("+s1+", "+s2+")")
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(ambiguousCoordinates("(123)"))
}
