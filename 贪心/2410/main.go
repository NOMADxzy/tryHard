package main

import (
	"fmt"
	"sort"
)

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Slice(players, func(i, j int) bool {
		return players[i] < players[j]
	})
	sort.Slice(trainers, func(i, j int) bool {
		return trainers[i] < trainers[j]
	})
	i, j := 0, 0
	m, n := len(players), len(trainers)
	ans := 0
	for i < m && j < n {
		for j < n && trainers[j] < players[i] {
			j++
		}
		if j == n {
			break
		}
		ans++
		i++
		j++
	}
	return ans
}

func main() {
	fmt.Println(matchPlayersAndTrainers([]int{4, 7, 9}, []int{8, 2, 5, 8}))
}
