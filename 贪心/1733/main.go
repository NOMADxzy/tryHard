package main

import "fmt"

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	alreadyLanguageCnt := make([]int, n)
	knowOnes := make([]map[int]bool, n)
	for i, vals := range languages {
		for _, val := range vals {
			if knowOnes[val-1] == nil {
				knowOnes[val-1] = map[int]bool{}
			}
			knowOnes[val-1][i] = true
		}
	}
	canCommunicate := func(a, b int) bool {
		for _, v := range languages[a-1] {
			if knowOnes[v-1][b-1] {
				return true
			}
		}
		return false
	}
	mark := make([]bool, len(languages))
	strangerCnts := 0
	for _, friendship := range friendships {
		a, b := friendship[0], friendship[1]
		if !canCommunicate(a, b) {
			mark[a-1] = true
			mark[b-1] = true
		}
	}
	for i := 0; i < len(languages); i++ {
		if mark[i] {
			for _, l := range languages[i] {
				alreadyLanguageCnt[l-1]++
			}
			strangerCnts++
		}
	}
	maxLanguage := 0
	for _, cnt := range alreadyLanguageCnt {
		maxLanguage = max(maxLanguage, cnt)
	}
	return strangerCnts - maxLanguage
}

func main() {
	fmt.Println(minimumTeachings(3, [][]int{{1, 3}, {2}, {1, 2}, {3}}, [][]int{{1, 2}, {1, 4}, {2, 3}, {3, 4}}))
}
