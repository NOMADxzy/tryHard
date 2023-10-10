package main

import "fmt"

func check(charCnts [][]int, left, right int) bool {
	for i := 0; i < len(charCnts); i++ {
		if left == 0 {
			if charCnts[i][right]%2 != 0 {
				return false
			}
		} else if (charCnts[i][right]-charCnts[i][left-1])%2 != 0 {
			return false
		}
	}
	return true
}

func findTheLongestSubstring(s string) int {
	charCnts := make([][]int, 5)
	for i := 0; i < 5; i++ {
		charCnts[i] = make([]int, len(s))
	}
	idxMap := map[byte]int{'a': 0, 'e': 1, 'i': 2, 'o': 3, 'u': 4}

	if s[0] == 'a' || s[0] == 'e' || s[0] == 'i' || s[0] == 'o' || s[0] == 'u' {
		charCnts[idxMap[s[0]]][0]++
	}
	for i := 1; i < len(s); i++ {
		for j := 0; j < 5; j++ {
			charCnts[j][i] = charCnts[j][i-1]
		}
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			idx := idxMap[s[i]]
			charCnts[idx][i]++
		}
	}

	maxLen := 0
	for i := 0; i < len(s); i++ {
		if check(charCnts, 0, i) && maxLen < i+1 {
			maxLen = i + 1
			continue
		}
		for j := 1; j <= i-maxLen; j++ {
			if check(charCnts, j, i) {
				if i-j+1 > maxLen {
					maxLen = i - j + 1
				}
				break
			}
		}
	}
	return maxLen
}

func main() {
	fmt.Println(findTheLongestSubstring("yopumzgd"))
}
