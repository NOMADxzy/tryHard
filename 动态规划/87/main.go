package main

import "fmt"

func isScramble(s1 string, s2 string) bool {
	m := len(s1)
	dp := make([][][]bool, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]bool, m)
		for j := 0; j < m; j++ {
			dp[i][j] = make([]bool, m)
			dp[i][j][0] = s1[i] == s2[j]
		}
	}
	for l := 1; l < m; l++ {
		for i := 0; i < m-l; i++ {
			for j := 0; j < m-l; j++ {
				left1 := i
				left2 := j
				for leftLen := 1; leftLen <= l; leftLen++ {
					rightLen := l + 1 - leftLen
					if dp[left1][left2][leftLen-1] && dp[left1+leftLen][left2+leftLen][rightLen-1] ||
						dp[left1][left2+rightLen][leftLen-1] && dp[left1+leftLen][left2][rightLen-1] {
						dp[left1][left2][l] = true
					}
				}
			}

		}
	}
	return dp[0][0][m-1]
}

// 递归法超时
func isScramble1(s1 string, s2 string) bool {
	m := len(s1)
	if m == 1 {
		return s1[0] == s2[0]
	}
	map1, map2 := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(s1); i++ {
		map1[s1[i]]++
		map2[s2[i]]++
	}
	for key, val := range map1 {
		if map2[key] != val {
			return false
		}
	}
	for leftLen := 1; leftLen < m; leftLen++ {
		if isScramble(s1[:leftLen], s2[:leftLen]) && isScramble(s1[leftLen:], s2[leftLen:]) ||
			isScramble(s1[:leftLen], s2[m-leftLen:]) && isScramble(s1[leftLen:], s2[:m-leftLen]) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(isScramble("abc", "bca"))
}
