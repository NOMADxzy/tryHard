package main

import (
	"fmt"
	"sort"
)

func dividePlayers(skill []int) int64 {
	sort.Slice(skill, func(i, j int) bool {
		return skill[i] < skill[j]
	})
	ans := int64(0)
	target := skill[0] + skill[len(skill)-1]
	for i := 0; i < len(skill)/2; i++ {
		if skill[i]+skill[len(skill)-1-i] != target {
			return int64(-1)
		}
		ans += int64(skill[i] * skill[len(skill)-1-i])
	}
	return ans
}

func main() {
	fmt.Println(dividePlayers([]int{3, 2, 5, 1, 3, 4}))
}
