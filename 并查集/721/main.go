package main

import (
	"fmt"
	"sort"
)

func accountsMerge(accounts [][]string) [][]string {
	emailIdx := map[string]int{}
	//var emailName []string
	parents := make([]int, len(accounts))
	for i := 0; i < len(parents); i++ {
		parents[i] = i
	}

	var find func([]int, int) int
	find = func(parents []int, index int) int {
		if parents[index] != index {
			return find(parents, parents[index])
		} else {
			return index
		}
	}
	union := func(parents []int, index1, index2 int) {
		parents[find(parents, index2)] = find(parents, index1)
	}

	//id := 0
	for idx, account := range accounts {
		for i := 1; i < len(account); i++ {
			if preIdx, ok := emailIdx[account[i]]; ok {
				union(parents, preIdx, idx)
				account[i] = ""
			} else {
				emailIdx[account[i]] = idx
			}
		}
	}

	var ans [][]string
	mailsUnion := map[int][]string{}

	for idx, account := range accounts {
		root := find(parents, idx)
		mailsUnion[root] = append(mailsUnion[root], account[1:]...)
	}

	for i, mails := range mailsUnion {
		sort.Strings(mails)
		start := 0
		for mails[start] == "" {
			start++
		}
		ans = append(ans, append([]string{accounts[i][0]}, mails[start:]...))
	}
	return ans
}

func main() {
	//accounts := [][]string{{"John", "johnsmith@mail.com", "john00@mail.com"}, {"John", "johnnybravo@mail.com"}, {"John", "johnsmith@mail.com", "john_newyork@mail.com"},
	//	{"Mary", "mary@mail.com"}}
	accounts := [][]string{{"David", "David0@m.co", "David1@m.co"}, {"David", "David3@m.co", "David4@m.co"}, {"David", "David4@m.co", "David5@m.co"},
		{"David", "David2@m.co", "David3@m.co"}, {"David", "David1@m.co", "David2@m.co"}}
	fmt.Println(accountsMerge(accounts))
}
