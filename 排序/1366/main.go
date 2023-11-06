package main

import (
	"fmt"
	"sort"
)

func rankTeams(votes []string) string {
	voteCnt := make([][]int, 26)
	for i := 0; i < 26; i++ {
		voteCnt[i] = make([]int, 26)
	}

	for _, vote := range votes {
		for r, c := range vote {
			voteCnt[c-'A'][r]++
		}
	}

	rank := make([]byte, len(votes[0]))
	for i := 0; i < len(votes[0]); i++ {
		rank[i] = votes[0][i]
	}

	sort.Slice(rank, func(i, j int) bool {
		size := len(votes[0])
		for k := 0; k < size; k++ {
			if voteCnt[rank[i]-'A'][k] > voteCnt[rank[j]-'A'][k] {
				return true
			} else if voteCnt[rank[i]-'A'][k] < voteCnt[rank[j]-'A'][k] {
				return false
			}
		}
		return rank[i] < rank[j]
	})
	return string(rank)
}

func main() {
	fmt.Println(rankTeams([]string{"ABC", "ACB", "ABC", "ACB", "ACB"}))
}
