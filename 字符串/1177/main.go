package main

import "fmt"

func canMakePaliQueries(s string, queries [][]int) []bool {
	charCnts := make([][]int, 26)
	for i := 0; i < 26; i++ {
		charCnts[i] = make([]int, len(s))
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		idx := int(c - 'a')
		if i == 0 {
			charCnts[idx][i]++
		} else {
			charCnts[idx][i] += charCnts[idx][i-1] + 1
			for j := 0; j < 26; j++ {
				if j != idx {
					charCnts[j][i] = charCnts[j][i-1]
				}
			}
		}
	}
	ans := make([]bool, len(queries))
	for i := 0; i < len(queries); i++ {
		q := queries[i]
		l, r := q[0], q[1]
		if l == r {
			ans[i] = true
			continue
		}
		samePairs := 0
		for j := 0; j < 26; j++ {
			if l == 0 {
				samePairs += charCnts[j][r] / 2
			} else {
				samePairs += (charCnts[j][r] - charCnts[j][l-1]) / 2
			}
		}
		need := (r - l + 1 - samePairs*2) / 2
		ans[i] = need <= q[2]
	}
	return ans
}

func main() {
	fmt.Println(canMakePaliQueries("abcda", [][]int{{3, 3, 0}, {1, 2, 0}, {0, 3, 1}, {0, 3, 2}, {0, 4, 1}}))
}
